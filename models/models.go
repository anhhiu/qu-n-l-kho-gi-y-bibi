package models

import "time"

// Supplier (Nhà cung cấp)
type Supplier struct {
	SupplierID   int    `gorm:"primaryKey;autoIncrement;column:supplier_id" json:"supplier_id"`
	SupplierName string `gorm:"size:50;column:supplier_name" json:"supplier_name"`
	Address      string `gorm:"size:100;column:address" json:"address"`
	Phone        string `gorm:"size:10;column:phone" json:"phone"`
	Email        string `gorm:"size:50;column:email" json:"email"`
	Website      string `gorm:"size:50;column:website" json:"website"`
}

// Customer (Khách hàng)
type Customer struct {
	CustomerID   int    `gorm:"primaryKey;autoIncrement;column:customer_id" json:"customer_id"`
	CustomerName string `gorm:"size:50;column:customer_name" json:"customer_name"`
	Email        string `gorm:"size:50;column:email" json:"email"`
	Phone        string `gorm:"size:10;column:phone" json:"phone"`
	Address      string `gorm:"size:100;column:address" json:"address"`
	Password     string `gorm:"size:50;column:password" json:"password"`
}

// Product (Sản phẩm)
type Product struct {
	ProductID   int     `gorm:"primaryKey;autoIncrement;column:product_id" json:"product_id"`
	ProductName string  `gorm:"size:50;column:product_name" json:"product_name"`
	Brand       string  `gorm:"size:50;column:brand" json:"brand"`
	Size        string  `gorm:"size:20;column:size" json:"size"`
	Color       string  `gorm:"size:50;column:color" json:"color"`
	Quantity    int     `gorm:"column:quantity" json:"quantity"`
	Price       float64 `gorm:"type:decimal(10,2);column:price" json:"price"`
	SupplierID  int     `gorm:"column:supplier_id" json:"supplier_id"`
	Image       string  `gorm:"size:255;column:image" json:"image"`
	Description string  `gorm:"type:text;column:description" json:"description"`

	Supplier *Supplier `gorm:"foreignKey:SupplierID" json:"supplier"`
}

/* // Order (Đơn hàng)
type Order struct {
	OrderID     int       `gorm:"primaryKey;autoIncrement;column:order_id" json:"order_id"`
	CustomerID  int       `gorm:"column:customer_id" json:"customer_id"`
	OrderDate   time.Time `gorm:"type:date;column:order_date" json:"order_date"`
	TotalAmount float64   `gorm:"type:decimal(10,2);column:total_amount" json:"total_amount"`
	Status      string    `gorm:"size:50;column:status" json:"status"`

	Customer *Customer `gorm:"foreignKey:CustomerID" json:"customer"`
}

// OrderDetail (Chi tiết đơn hàng)
type OrderDetail struct {
	OrderDetailID int     `gorm:"primaryKey;autoIncrement;column:order_detail_id" json:"order_detail_id"`
	OrderID       int     `gorm:"column:order_id" json:"order_id"`
	ProductID     int     `gorm:"column:product_id" json:"product_id"`
	Quantity      int     `gorm:"column:quantity" json:"quantity"`
	Price         float64 `gorm:"type:decimal(10,2);column:price" json:"price"`

	Order   *Order   `gorm:"foreignKey:OrderID" json:"order"`
	Product *Product `gorm:"foreignKey:ProductID" json:"product"`
} */
// Order (Đơn hàng)
type Order struct {
	OrderID     int       `gorm:"primaryKey;autoIncrement;column:order_id" json:"order_id"`
	CustomerID  int       `gorm:"column:customer_id" json:"customer_id"`
	OrderDate   time.Time `gorm:"type:date;column:order_date" json:"order_date"`
	TotalAmount float64   `gorm:"type:decimal(10,2);column:total_amount" json:"total_amount"`
	Status      string    `gorm:"size:50;column:status" json:"status"`

	Customer *Customer `gorm:"foreignKey:CustomerID" json:"customer"`
}

// OrderDetail (Chi tiết đơn hàng)
type OrderDetail struct {
	OrderDetailID int     `gorm:"primaryKey;autoIncrement;column:order_detail_id" json:"order_detail_id"`
	OrderID       int     `gorm:"column:order_id" json:"order_id"`
	ProductID     int     `gorm:"column:product_id" json:"product_id"`
	Quantity      int     `gorm:"column:quantity" json:"quantity"`
	UnitPrice     float64 `gorm:"type:decimal(10,2);column:unit_price" json:"unit_price"`

	Order   *Order   `gorm:"foreignKey:OrderID" json:"order"`
	Product *Product `gorm:"foreignKey:ProductID" json:"product"`
}

type Users struct {
	Id       int    `gorm:"primaryKey;autoInCrement;user_id" json:"user_id"`
	Role     string `gorm:"column:role" json:"role"`
	UserName string `gorm:"column:username" json:"username"`
	PassWord string `gorm:"column:password" json:"password"`
}
