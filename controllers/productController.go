package controllers

import (
	"bibi/config"
	"bibi/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @tags products
// @summary Get all products
// @router /product/ [get]
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// @tags products
// @summary Create product
// @param product body models.Product true "Product data"
// @router /product/ [post]
func CreateProduct(c *gin.Context) {
	var input struct {
		ProductName string           `json:"product_name"`
		Brand       string           `json:"brand"`
		Size        string           `json:"size"`
		Color       string           `json:"color"`
		Quantity    int              `json:"quantity"`
		Price       float64          `json:"price"`
		SupplierID  int              `json:"supplier_id"`
		Image       string           `json:"image"`
		Description string           `json:"description"`
		Supplier    *models.Supplier `json:"supplier"` // Nếu cần thiết
	}

	// Kiểm tra lỗi JSON đầu vào
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Trường hợp khi nhà cung cấp chưa được tạo
	if input.SupplierID == 0 && input.Supplier != nil {
		config.DB.Create(&input.Supplier)
		input.SupplierID = input.Supplier.SupplierID
	}

	product := models.Product{
		ProductName: input.ProductName,
		Brand:       input.Brand,
		Size:        input.Size,
		Color:       input.Color,
		Quantity:    input.Quantity,
		Price:       input.Price,
		SupplierID:  input.SupplierID,
		Image:       input.Image,
		Description: input.Description,
	}

	config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// @tags products
// @summary Update product by id
// @param product_id path int true "ProductID"
// @param product body models.Product true "Products info"
// @router /product/{product_id} [put]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.Where("product_id = ?", c.Param("product_id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cập nhật thông tin sản phẩm
	product.ProductName = input.ProductName
	product.Brand = input.Brand
	product.Size = input.Size
	product.Color = input.Color
	product.Quantity = input.Quantity
	product.Price = input.Price
	product.SupplierID = input.SupplierID
	product.Image = input.Image
	product.Description = input.Description

	config.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// @tags products
// @summary Delete product by id
// @param product_id path int true "ProductID"
// @router /product/{product_id} [delete]
func DeleteProduct(c *gin.Context) {
	if err := config.DB.Delete(&models.Product{}, c.Param("product_id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// @tags products
// @summary Get product by id
// @param product_id path int true "ProductID"
// @router /product/{product_id} [get]
func GetProductByID(c *gin.Context) {
	productIDStr := c.Param("product_id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	var product models.Product
	if err := config.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
		log.Printf("Error retrieving product: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// quản lý tồn kho
func UpdateProductByIdQuantity(c *gin.Context) {
	var product models.Product
	if err := config.DB.Where("product_id = ?", c.Param("product_id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Nhận dữ liệu từ body yêu cầu
	var input struct {
		QuantityChange int `json:"quantity_change"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cập nhật số lượng mới
	newQuantity := product.Quantity + input.QuantityChange

	// Kiểm tra số lượng không được âm
	if newQuantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Lưu số lượng mới vào database
	product.Quantity = newQuantity
	config.DB.Save(&product)

	// Phản hồi kết quả
	c.JSON(http.StatusOK, gin.H{
		"message":      "Stock quantity updated successfully",
		"new_quantity": product.Quantity,
	})
}

// lấy tất cả các sản phẩm  theo id khách hàng đã mua : chưa chạy
func GetPurchasedProductsByCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")

	var products []models.Product
	if err := config.DB.Table("order_details").
		Select("products.*, order_details.quantity").
		Joins("JOIN orders ON orders.order_id = order_details.order_id").
		Joins("JOIN products ON products.product_id = order_details.product_id").
		Where("orders.customer_id = ?", customerID).
		Scan(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purchased products"})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found for this customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// báo cáo tồn kho

func GetInventoryReport1(c *gin.Context) {
	var products []models.Product

	// Thực hiện truy vấn để lấy tất cả sản phẩm và số lượng tồn kho
	if err := config.DB.Find(&products).Error; err != nil {
		log.Printf("Error retrieving products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory"})
		return
	}

	// Tạo danh sách báo cáo tồn kho
	inventoryReport := make([]gin.H, len(products))

	for i, product := range products {
		inventoryReport[i] = gin.H{
			"product_id":   product.ProductID,
			"product_name": product.ProductName,
			"quantity":     product.Quantity,
			"price":        product.Price,
			"brand":        product.Brand,
			"size":         product.Size,
			"color":        product.Color,
		}
	}

	// Trả về kết quả
	c.JSON(http.StatusOK, gin.H{"inventory": inventoryReport})
}
