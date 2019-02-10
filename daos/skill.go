package daos

import "freelancers/models"

type SkillDAO struct{}

func NewSkillDAO() *SkillDAO {
	return &SkillDAO{}
}

func (dao *SkillDAO) findSkills(query string) map[int]models.Skill {

	return nil
}
