package daos

import (
	"freelancers/app"
	"freelancers/models"
)

type ProjectDAO struct{}

func NewProjectDAO() *ProjectDAO {
	return &ProjectDAO{}
}

func (dao *ProjectDAO) CreateProject(rs app.RequestScope, project *models.Project) error {
	create := rs.Db().Create(&project)
	return create.Error
}

func (dao *ProjectDAO) GetProjectByID(rs app.RequestScope, id uint) (project models.Project) {
	rs.Db().Where("id = ?", id).Preload("Skills").Preload("User").First(&project)
	return
}

func (dao *ProjectDAO) GetAllProjects(rs app.RequestScope) (err error, projects []models.Project) {
	db := rs.Db().Model(rs.User()).Related(&projects)

	return db.Error, projects
}
