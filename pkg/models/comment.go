package models

import "time"

type Comment struct {
	Id           int       `json:"id"`
	UserId       int       `json:"userId"`
	ProductId    int       `json:"productId"`
	CreationDate time.Time `json:"creationDate"`
	Rate         int       `json:"rate"`
	Content      string    `json:"content"`
}

func (Comment) TableName() string {
	return "t_comment"
}
