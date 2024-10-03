// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customer/": {
            "get": {
                "description": "Retrieve a list of customers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Get all customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Customer"
                            }
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "customers"
                ],
                "summary": "create customer",
                "parameters": [
                    {
                        "description": "customer data",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/customer/{customer_id}": {
            "get": {
                "tags": [
                    "customers"
                ],
                "summary": "get customer by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "customer_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "tags": [
                    "customers"
                ],
                "summary": "update customer by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "customer_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer info",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "tags": [
                    "customers"
                ],
                "summary": "delete customer by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "customer_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/login/": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "đăng nhập tài khoản",
                "parameters": [
                    {
                        "description": "Users data",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/order/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get all orders",
                "responses": {}
            }
        },
        "/order/{order_id}": {
            "put": {
                "tags": [
                    "orders"
                ],
                "summary": "update order by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated order data",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "tags": [
                    "orders"
                ],
                "summary": "delele order by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "OrderID",
                        "name": "orders_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/orderdetail/": {
            "get": {
                "tags": [
                    "orderdetails"
                ],
                "summary": "get all orderdetails",
                "responses": {}
            }
        },
        "/orderdetail/{order_detail_id}": {
            "get": {
                "description": "Retrieve a specific order detail by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orderdetails"
                ],
                "summary": "Get order detail by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "OrderDetailID",
                        "name": "order_detail_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/product/": {
            "get": {
                "tags": [
                    "products"
                ],
                "summary": "get allproducts",
                "responses": {}
            },
            "post": {
                "tags": [
                    "products"
                ],
                "summary": "create product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/product/{product_id}": {
            "get": {
                "tags": [
                    "products"
                ],
                "summary": "get product by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "tags": [
                    "products"
                ],
                "summary": "update product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Products info",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "tags": [
                    "products"
                ],
                "summary": "delete product by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/register/": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "đăng kí tài khoản",
                "parameters": [
                    {
                        "description": "Users data",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/report/inventory": {
            "get": {
                "description": "Lấy báo cáo tồn kho",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report"
                ],
                "summary": "Thống kê báo cáo tồn kho",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/report/order": {
            "get": {
                "description": "Lấy báo cáo đơn hàng: xử lý, đang chờ, và bị hủy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report"
                ],
                "summary": "Thống kê đơn hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/report/revenue": {
            "get": {
                "description": "Lấy báo cáo doanh thu theo tháng",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report"
                ],
                "summary": "Thống kê doanh thu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/supplier/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "suppliers"
                ],
                "summary": "Get all suppliers",
                "responses": {}
            },
            "post": {
                "tags": [
                    "suppliers"
                ],
                "summary": "delete supplier",
                "parameters": [
                    {
                        "description": "Supplier data",
                        "name": "supplier",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Supplier"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/supplier/{supplier_id}": {
            "get": {
                "tags": [
                    "suppliers"
                ],
                "summary": "Get supplier by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Supplier ID",
                        "name": "supplier_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "tags": [
                    "suppliers"
                ],
                "summary": "update supplier",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Supplier ID",
                        "name": "supplier_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Supplier info",
                        "name": "supplier",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Supplier"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "tags": [
                    "suppliers"
                ],
                "summary": "delete supplier",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Supplier ID",
                        "name": "supplier_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Customer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "customer_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "size": {
                    "type": "string"
                },
                "supplier": {
                    "$ref": "#/definitions/models.Supplier"
                },
                "supplier_id": {
                    "type": "integer"
                }
            }
        },
        "models.Supplier": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "supplier_id": {
                    "type": "integer"
                },
                "supplier_name": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
