package models

import (
	"time"
)

type Product struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ShopId         int    `json:"shopId"`
	CurrentPriceId int    `json:"currentPriceId"`
	PriorityValue  int    `json:"priorityValue"`
}

type ProductPrice struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	ProductId    int       `json:"productId"`
	Price        float32   `json:"price"`
	Discount     float32   `json:"discount"`
	CreationDate time.Time `json:"creationDate"`
}

type ProductCharacteristic struct {
	Id        int    `json:"id"`
	ProductId int    `json:"productId"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

type ProductToOrder struct {
	Id        int `json:"id"`
	OrderId   int `json:"orderId"`
	ProductId int `json:"productId"`
	Amount    int `json:"amount"`
}

type ProductToCart struct {
	Id        int       `json:"id"`
	ProductId int       `json:"productId"`
	UserId    int       `json:"userId"`
	AddDate   time.Time `json:"addDate"`
	Amount    int       `json:"amount"`
}

type ProductToFavourite struct {
	Id        int       `json:"id"`
	ProductId int       `json:"productId"`
	UserId    int       `json:"userId"`
	AddDate   time.Time `json:"addDate"`
}

func (Product) TableName() string {
	return "t_product"
}

func (ProductPrice) TableName() string {
	return "t_product_price"
}

func (ProductCharacteristic) TableName() string {
	return "t_product_characteristic"
}

func (ProductToOrder) TableName() string {
	return "t_product_to_order"
}

func (ProductToCart) TableName() string {
	return "t_product_to_cart"
}

func (ProductToFavourite) TableName() string {
	return "t_product_to_favourites"
}
