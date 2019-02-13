package models

import "github.com/jinzhu/gorm"

type Skill struct {
	gorm.Model
	Name string
}

func (s *Skill) GetID() uint {
	return s.ID
}

func (s *Skill) GetName() string {
	return s.Name
}
