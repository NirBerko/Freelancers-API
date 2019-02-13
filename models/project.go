package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Title       string
	Description string
	Skills      []Skill `gorm:"many2many:project_skills;"`
	BudgetType  uint
	BudgetLevel uint
}

func (p *Project) GetID() uint {
	return p.ID
}

func (p *Project) GetTitle() string {
	return p.Title
}

func (p *Project) GetDescription() string {
	return p.Description
}

func (p *Project) GetSkills() []Skill {
	return p.Skills
}

func (p *Project) GetBudgetType() uint {
	return p.BudgetType
}

func (p *Project) GetBudgetLevel() uint {
	return p.BudgetLevel
}
