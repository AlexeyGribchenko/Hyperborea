package models

type Shop struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (Shop) TableName() string {
	return "t_shop"
}
