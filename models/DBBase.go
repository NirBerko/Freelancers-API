package models

import "time"

type DBBase struct {
	ID        uint       `gorm:"primary_key" `
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
