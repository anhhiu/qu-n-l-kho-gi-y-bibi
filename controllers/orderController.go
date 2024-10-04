package controllers

import (
	"bibi/config"
	"bibi/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get all orders
// @Tags orders
// @Accept json
// @Produce json
// @Router /order/ [get]
func GetAllOrder(c *gin.Context) {
	var orders []models.Order

	// Lấy tất cả các đơn hàng từ cơ sở dữ liệu
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	// Kiểm tra xem có đơn hàng nào không
	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// @Summary Create order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body object true "Orders data"
// @Router /order/ [post]

func CreateOrder(c *gin.Context) {
	// Khai báo cấu trúc đầu vào
	var input struct {
		CustomerID int       `json:"customer_id"`
		OrderDate  time.Time `json:"order_date"`
		Products   []struct {
			ProductID int     `json:"product_id"`
			Quantity  int     `json:"quantity"`
			UnitPrice float64 `json:"unit_price"`
		} `json:"products"`
	}

	// Kiểm tra lỗi JSON đầu vào
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Khởi tạo transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	// Tạo đơn hàng mới
	order := models.Order{
		CustomerID: input.CustomerID,
		OrderDate:  time.Now(),
		Status:     "Đã Mua", // Trạng thái mặc định
	}

	// Lưu đơn hàng vào cơ sở dữ liệu
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback() // Rollback transaction nếu có lỗi
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Tính toán tổng tiền
	var totalAmount float64
	for _, p := range input.Products {
		// Tính toán tổng tiền từ sản phẩm
		totalAmount += float64(p.Quantity) * p.UnitPrice

		// Lấy sản phẩm từ cơ sở dữ liệu
		var product models.Product
		if err := tx.First(&product, p.ProductID).Error; err != nil {
			tx.Rollback() // Rollback transaction nếu có lỗi
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}

		// Kiểm tra xem số lượng trong kho có đủ không
		if product.Quantity < p.Quantity {
			tx.Rollback() // Rollback transaction nếu có lỗi
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product ID: " + strconv.Itoa(p.ProductID)})
			return
		}

		// Tạo chi tiết đơn hàng
		orderDetail := models.OrderDetail{
			OrderID:   order.OrderID,
			ProductID: p.ProductID,
			Quantity:  p.Quantity,
			UnitPrice: p.UnitPrice,
		}

		// Lưu chi tiết đơn hàng vào cơ sở dữ liệu
		if err := tx.Create(&orderDetail).Error; err != nil {
			tx.Rollback() // Rollback transaction nếu có lỗi
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order detail"})
			return
		}

		// Giảm số lượng sản phẩm trong kho
		product.Quantity -= p.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback() // Rollback transaction nếu có lỗi
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
			return
		}
	}

	// Cập nhật tổng tiền cho đơn hàng
	order.TotalAmount = totalAmount
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback() // Rollback transaction nếu có lỗi
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order total amount"})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// Trả về phản hồi
	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

// @Summary Update order by id
// @Tags orders
// @Param order_id path int true "Order ID"
// @Param order body object true "Updated order data"
// @Router /order/{order_id} [put]
func UpdateOrderById(c *gin.Context) {
	var input struct {
		CustomerID int       `json:"customer_id"`
		OrderDate  time.Time `json:"order_date"`
		Products   []struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		} `json:"products"`
	}

	// kiểm tra json đầu vào
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// tìm kiếm đơn hàng theo iid
	var order models.Order
	if err := config.DB.First(&order, c.Param("order_id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Cập nhật các trường lệnh
	order.CustomerID = input.CustomerID
	order.OrderDate = time.Now()

	// cập nhật đơn hàng chi tiết
	for _, p := range input.Products {
		// Get the existing product
		var product models.Product
		if err := config.DB.First(&product, p.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}

		// cập nhật sản phẩm trong kho
		product.Quantity -= p.Quantity // Adjust quantity
		if product.Quantity < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product ID: " + strconv.Itoa(p.ProductID)})
			return
		}

		// cập nhật và lưu vào product
		if err := config.DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
			return
		}
	}

	// cập nhật và lưu order
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": order})
}

// @Summary Delete order by id
// @Tags orders
// @Param order_id path int true "Order ID"
// @Router /order/{order_id} [delete]
func DeleteOrderById(c *gin.Context) {
	// tìm hóa đơn theo id
	var order models.Order
	if err := config.DB.First(&order, c.Param("order_id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// tìm hóa đơn chi tiết
	var orderDetails []models.OrderDetail
	if err := config.DB.Where("order_id = ?", order.OrderID).Find(&orderDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order details"})
		return
	}

	// khôi phục lại số lượng sản phẩm trong kho
	for _, detail := range orderDetails {
		var product models.Product
		if err := config.DB.First(&product, detail.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}

		product.Quantity += detail.Quantity

		// cập nhật và lưu vào  product
		if err := config.DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
			return
		}
	}

	// xóa thông tin chi tiết về đơn hàng
	if err := config.DB.Where("order_id = ?", order.OrderID).Delete(&models.OrderDetail{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order details"})
		return
	}

	// xóa đơn hàng
	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// @tags orders
// @summary Get order by id
// @param orders_id path int true "OrderID"
// @router /order/{order_id} [get]
func GetOrderById(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("order_id = ?", c.Param("order_id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}
