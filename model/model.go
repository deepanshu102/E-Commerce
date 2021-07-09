package model

import (
	"gorm.io/gorm"
)

type (
	Category struct {
		gorm.Model
		Name        string    `json:"name"`
		ProductList []Product `json:"productList"`
	}
	Product struct {
		gorm.Model
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		Stock       int     `json:"stock"`
		CategoryID  int
		Category    *Category `json:"category"`
	}

	Cart struct {
		gorm.Model
		ProductId int      `json:"productId"`
		Product   *Product `json:"product"`
		Quantity  int      `json:"quantity"`
		Amount    float64  `json:"amount"`
		UsersID   int
		User      *Users `gorm:"foreignKey:UsersID"`
	}
	Users struct {
		gorm.Model
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Role      string `json:"role"`
		CartList  []Cart
		OrderList []Orders
	}

	ProductPurchaseStock struct {
		gorm.Model
		ProductId int
		Product   Product `json:"product"`
		Quantity  int     `json:"quantity"`
		Amount    float64 `json:"amount"`
		OrdersID  int
		Order     *Orders `gorm:"foreignKey:OrdersID"`
	}

	Address struct {
		gorm.Model
		FirstAddress  string `json:"firstAddress"`
		SecondAddress string `json:"secondAddress"`
		State         string `json:"state"`
		Country       string `json:"country"`
		Pincode       string `json:"pincode"`
		ShippingID    int
		Shipping      *Shipping `gorm:"foreignKey:ShippingID"`
	}
	Shipping struct {
		gorm.Model

		Status   string   `json:"status"`
		Address  *Address `json:"shippingaddress"`
		OrdersID int
		Order    *Orders `gorm:"foreignKey:OrdersID"`
	}
	Orders struct {
		gorm.Model

		Status      string                 `json:"status"`
		ProductList []ProductPurchaseStock `json:"productlist"`
		Amount      float64                `json:"amount"`
		Shipping    *Shipping              `json:"shippingAddress"`
		UsersID     int
		User        *Users `gorm:"foreignKey:UsersID"`
	}
)
