package models

import "github.com/lib/pq"

type Project struct {
	DBBase
	Title       string
	Description string
	Skills      pq.StringArray `gorm:"type:varchar(64)[]"`
}
