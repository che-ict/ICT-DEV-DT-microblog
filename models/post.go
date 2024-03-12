package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content string
	UserId  uint
	User    User
}

type User struct {
	gorm.Model
	Username    string `gorm:"unique"`
	Password    string
	DisplayName string
	Posts       []Post
}
