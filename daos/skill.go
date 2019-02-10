package daos

import (
	"freelancers/app"
	"freelancers/models"
)

type SkillDAO struct{}

func NewSkillDAO() *SkillDAO {
	return &SkillDAO{}
}

func (dao *SkillDAO) findSkills(rs app.RequestScope, query string) (skills map[int]models.Skill) {
	rs.Db().Where("name LIKE ?", "%"+query+"").Find(&skills)
	return nil
}
