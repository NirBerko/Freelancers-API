package daos

import (
	"fmt"
	"freelancers/app"
	"freelancers/models"
)

type SkillDAO struct{}

func NewSkillDAO() *SkillDAO {
	return &SkillDAO{}
}

func (dao *SkillDAO) CreateSkill(rs app.RequestScope, skill *models.Skill) {
	var findSkill models.Skill
	rs.Db().Where("name LIKE ?", skill.Name).Find(&findSkill)

	fmt.Println(findSkill)

	//rs.Db().Create(&skill)
}

/*
func (internalModels *SkillDAO) findSkills(rs app.RequestScope, query string) (skills map[int]models.Skill) {
	rs.Db().Where("name LIKE ?", "%"+query+"").Find(&skills)
	return nil
}
*/
