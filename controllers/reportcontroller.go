package controllers

import (
	"bibi/config"
	"bibi/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetInventoryReport godoc
// @Summary Thống kê báo cáo tồn kho
// @Description Lấy báo cáo tồn kho
// @Tags report
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Router /report/inventory [get]
func GetInventoryReport(c *gin.Context) {
	var products []models.Product
	var report models.InventoryReport

	// Lấy tất cả sản phẩm từ cơ sở dữ liệu
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tính toán tổng số sản phẩm, tổng số lượng và tổng giá trị
	for _, product := range products {
		report.TotalProducts++
		report.TotalQuantity += product.Quantity
		report.TotalValue += float64(product.Quantity) * product.Price

		//  tồn kho thấp là nhỏ hơn hoặc bằng 5 sản phẩm
		if product.Quantity <= 5 {
			report.LowStockProducts = append(report.LowStockProducts, product)
		}
	}

	c.JSON(http.StatusOK, report)
}

// GetRevenueReport godoc
// @Summary Thống kê doanh thu
// @Description Lấy báo cáo doanh thu theo tháng
// @Tags report
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Router /report/revenue [get]
func GetRevenueReport(c *gin.Context) {
	var orders []models.Order
	revenueMap := make(map[string]float64) // để lưu trữ doanh thu theo tháng

	// Lấy tất cả đơn hàng từ cơ sở dữ liệu
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tính toán doanh thu theo tháng
	for _, order := range orders {
		month := order.OrderDate.Month() // Lấy tháng
		year := order.OrderDate.Year()   // Lấy năm
		key := fmt.Sprintf("%d-%d", year, month)

		revenueMap[key] += order.TotalAmount // Cộng dồn doanh thu cho tháng
	}

	var report []models.RevenueReport
	// Chuyển đổi map thành slice
	for key, revenue := range revenueMap {
		parts := strings.Split(key, "-")
		year, _ := strconv.Atoi(parts[0])
		month, _ := strconv.Atoi(parts[1])

		report = append(report, models.RevenueReport{
			Year:    year,
			Month:   month,
			Revenue: revenue,
		})
	}

	c.JSON(http.StatusOK, report)
}

// GetOrderReport godoc
// @Summary Thống kê đơn hàng
// @Param Authorization header string true "Bearer token"
// @Description Lấy báo cáo đơn hàng: xử lý, đang chờ, và bị hủy
// @Tags report
// @Accept json
// @Produce json
// @Router /report/order [get]
func GetOrderReport(c *gin.Context) {
	var orders []models.Order
	orderReports := make(map[string]*models.OrderReport) // để lưu trữ báo cáo theo trạng thái

	// Lấy tất cả đơn hàng từ cơ sở dữ liệu
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tính toán số lượng và doanh thu theo trạng thái
	for _, order := range orders {
		if _, exists := orderReports[order.Status]; !exists {
			orderReports[order.Status] = &models.OrderReport{
				Status:       order.Status,
				TotalOrders:  0,
				TotalRevenue: 0,
			}
		}
		orderReports[order.Status].TotalOrders++
		orderReports[order.Status].TotalRevenue += order.TotalAmount
	}

	var report []models.OrderReport
	// Chuyển đổi map thành slice
	for _, reportItem := range orderReports {
		report = append(report, *reportItem)
	}

	c.JSON(http.StatusOK, report)
}
