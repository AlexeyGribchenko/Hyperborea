package models

type Category struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parentId"`
	Name     string `json:"name"`
}

type CategoryToProduct struct {
	Id         int `json:"id"`
	ProductId  int `json:"productId"`
	CategoryId int `json:"categoryId"`
}

func (Category) TableName() string {
	return "t_category"
}

func (CategoryToProduct) TableName() string {
	return "t_category_to_product"
}
