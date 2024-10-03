package routes

import (
	"bibi/auth"
	"bibi/controllers"
	"bibi/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"                   // Swagger embed files
	httpSwagger "github.com/swaggo/http-swagger" // Http Swagger middleware
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: true, // Cho phép tất cả các nguồn gốc
		// Hoặc bạn có thể chỉ định các nguồn gốc cụ thể:
		// AllowOrigins: []string{"http://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 60 * 60, // 12 giờ
	}
	router.Use(cors.New(corsConfig))
	docs.SwaggerInfo.BasePath = "/api"

	// Cập nhật dòng này
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

	router.POST("/api/register", auth.Register)
	router.POST("/api/login", auth.Login)

	report := router.Group("/api/report")
	{
		report.Use(auth.AuthMiddleware("admin27"))
		// token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInJvbGUiOiJhZG1pbjI3In0.uOYO9gH59494Wvm38NYSdTO10FgiWaOw28WIX9mlPyY
		// Bao cao ton kho
		report.GET("/inventory", controllers.GetInventoryReport)
		// Bao cao doanh thu
		report.GET("/revenue", controllers.GetRevenueReport)
		// Bao cao đơn hàng
		report.GET("/order", controllers.GetOrderReport)
	}

	supplier := router.Group("/api/supplier")
	{
		supplier.GET("/", controllers.GetSupplier)
		supplier.GET("/:supplier_id", controllers.GetSupplierById)
		supplier.POST("/", controllers.CreateSupplier)
		supplier.PUT("/:supplier_id", controllers.UpdateSuplierById)
		supplier.DELETE("/:supplier_id", controllers.DeleteSupplierById)
	}

	customer := router.Group("/api/customer")
	{
		customer.GET("/", controllers.GetCustomer)
		customer.GET("/:customer_id", controllers.GetCustomerById)
		customer.POST("/", controllers.CreateCustomer)
		customer.PUT("/:customer_id", controllers.UpdateCustomerById)
		customer.DELETE("/:customer_id", controllers.DeleteCustomerById)
	}

	user := router.Group("/api/user")
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
	}

	product := router.Group("/api/product")
	{
		product.GET("/", controllers.GetAllProducts)
		product.POST("/", controllers.CreateProduct)
		product.GET("/:product_id/products", controllers.GetPurchasedProductsByCustomer)
		product.GET("/:product_id", controllers.GetProductByID)
		product.PUT("/:product_id", controllers.UpdateProduct)
		product.PUT("/:product_id/quantity", controllers.UpdateProductByIdQuantity)
		product.DELETE("/:product_id", controllers.DeleteProduct)
		product.GET("/inventory", controllers.GetInventoryReport)
	}

	order := router.Group("/api/order")
	{
		order.GET("/", controllers.GetAllOrder)
		order.GET("/:order_id", controllers.GetOrderById)
		order.POST("/", controllers.CreateOrder)
		order.PUT("/:order_id", controllers.UpdateOrderById)
		order.DELETE("/:order_id", controllers.DeleteOrderById)
	}

	orderdetail := router.Group("/api/orderdetail")
	{
		orderdetail.GET("/", controllers.GetALLOrderDetail)
		orderdetail.GET("/:order_detail_id", controllers.GetOrderDetailById)
	}

	return router
}
