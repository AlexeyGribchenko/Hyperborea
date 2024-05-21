package models

type User struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RoleToUser struct {
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}

func (User) TableName() string {
	return "t_user"
}

func (Role) TableName() string {
	return "t_role"
}

func (RoleToUser) TableName() string {
	return "t_role_to_user"
}
