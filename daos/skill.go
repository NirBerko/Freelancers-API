package daos

type SkillDAO struct{}

func NewSkillDAO() *SkillDAO {
	return &SkillDAO{}
}

/*
func (internalModels *SkillDAO) findSkills(rs app.RequestScope, query string) (skills map[int]models.Skill) {
	rs.Db().Where("name LIKE ?", "%"+query+"").Find(&skills)
	return nil
}
*/
