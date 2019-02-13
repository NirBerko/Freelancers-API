package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Project struct {
	gorm.Model
	Title       string
	Description string
	Skills      pq.StringArray `gorm:"type:varchar(64)[]"`
	BudgetType  int
	BudgetLevel int
}
