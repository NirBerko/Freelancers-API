package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"-"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
}

func (u User) GetID() uint {
	return u.ID
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetFirstName() string {
	return u.FirstName
}

func (u User) GetLastName() string {
	return u.LastName
}
