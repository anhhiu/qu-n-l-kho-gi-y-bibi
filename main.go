package main

import (
	"bibi/config"
	_ "bibi/docs"
	"bibi/models"
	"bibi/routes"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Đây là bản demo quản lý kho giày bóng đá bibsport!")

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Supplier{})
	config.DB.AutoMigrate(&models.Customer{})
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.Order{})
	config.DB.AutoMigrate(&models.OrderDetail{})
	config.DB.AutoMigrate(&models.Users{})
	fmt.Println("ánh xạ thành công")

	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Không thể khởi chạy server: %v", err)
	}
	fmt.Println("routes ok")

}
