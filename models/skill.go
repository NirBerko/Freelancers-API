package models

import "github.com/jinzhu/gorm"

type Skill struct {
	gorm.Model
	Name string
}
