package models

type Image struct {
	Id   int    `json:"id"`
	Path string `json:"path"`
}

type ImageToComment struct {
	Id        int `json:"id"`
	ImageId   int `json:"imageId"`
	CommentId int `json:"commentId"`
}

type ImageToProduct struct {
	Id        int `json:"id"`
	ProductId int `json:"productId"`
	ImageId   int `json:"imageId"`
}

func (Image) TableName() string {
	return "t_image"
}

func (ImageToComment) TableName() string {
	return "t_image_to_product"
}

func (ImageToProduct) TableName() string {
	return "t_image_to_product"
}
