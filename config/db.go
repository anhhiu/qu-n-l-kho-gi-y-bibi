package config

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	/* //dsn := "sqlserver://sa:hao123@FF01:1433?database=demobibi&instance=HAHAO"
	dsn := "sqlserver://sa:hao123@LAPTOP-7CAHEI3Q:1433?database=demobibi&instance=HATHANHHAO"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("ket noi khong thanh cong", err)
	}
	fmt.Println("con nect thanh cong")
	DB = database */
	// Tải tệp .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Không thể tải tệp .env")
    }

    // Lấy thông tin từ biến môi trường
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    dbInstance := os.Getenv("DB_INSTANCE")

    // Kiểm tra các giá trị để đảm bảo rằng chúng không bị trống
    if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" || dbInstance == "" {
        log.Fatal("Thiếu thông tin kết nối cơ sở dữ liệu")
    }

    // Tạo chuỗi kết nối với SQL Server
    dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&instance=%s",
        dbUser, dbPassword, dbHost, dbPort, dbName, dbInstance)

    // Mở kết nối đến cơ sở dữ liệu
    database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Kết nối không thành công:", err)
    }
    fmt.Println("Kết nối thành công")
    DB = database
}
