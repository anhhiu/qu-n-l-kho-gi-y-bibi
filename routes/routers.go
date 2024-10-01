package routes

import (
	"bibi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	supplier := router.Group("/supplier")
	{
		supplier.GET("/", controllers.GetSupplier)
		supplier.GET("/:supplier_id", controllers.GetSupplierById)
		supplier.POST("/", controllers.CreateSupplier)
		supplier.PUT("/:supplier_id", controllers.UpdateSuplierById)
		supplier.DELETE("/:supplier_id", controllers.DeleteSupplierById)
	}

	customer := router.Group("/customer")
	{
		customer.GET("/", controllers.GetCustomer)
		customer.GET("/:customer_id", controllers.GetCustomerById)
		customer.POST("/", controllers.CreateCustomer)
		customer.PUT("/:customer_id", controllers.UpdateCustomerById)
		customer.DELETE("/:customer_id", controllers.DeleteCustomerById)
	}

	user := router.Group("/user")
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
	}

	product := router.Group("/product")
	{
		product.GET("/", controllers.GetAllProducts)
		product.POST("/", controllers.CreateProduct)
		product.GET("/:product_id/products", controllers.GetPurchasedProductsByCustomer)
		product.GET("/:product_id", controllers.GetProductByID)
		product.PUT("/:product_id", controllers.UpdateProduct)
		product.PUT("/:product_id/quantity", controllers.UpdateProductByIdQuantity)
		product.DELETE("/product_id", controllers.DeleteProduct)
		product.GET("/inventory ", controllers.GetInventoryReport)
	}

	order := router.Group("/order")
	{
		order.GET("/", controllers.GetAllOrder)
		order.GET("/:order_id", controllers.GetOrderById)
		order.POST("/", controllers.CreateOrder)
		order.PUT("/:order_id", controllers.UpdateOrderById)
		order.DELETE("/:order_id", controllers.DeleteOrderById)
	}

	orderdetail := router.Group("/orderdetail")
	{
		orderdetail.GET("/", controllers.GetALLOrderDetail)
		orderdetail.GET("/:order_detail_id", controllers.GetOrderDetailById)
	}

	return router

}