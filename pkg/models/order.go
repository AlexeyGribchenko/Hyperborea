package models

import "time"

type Order struct {
	Id           int       `json:"id"`
	AddressId    int       `json:"addressId"`
	TotalPrice   float64   `json:"totalPrice"`
	CreationDate time.Time `json:"creationDate"`
	StatusId     int       `json:"statusId"`
	UserId       int       `json:"userId"`
}

type PickUpAddress struct {
	Id     int    `json:"id"`
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
}

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Order) TableName() string {
	return "t_order"
}

func (PickUpAddress) TableName() string {
	return "t_pick_up_address"
}

func (Status) TableName() string {
	return "t_status"
}
