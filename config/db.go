package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "sqlserver://sa:hao123@LAPTOP-7CAHEI3Q:1433?database=demobibi&instance=HATHANHHAO"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("ket noi khong thanh cong", err)
	}
	fmt.Println("con nect thanh cong")
	DB = database
}
