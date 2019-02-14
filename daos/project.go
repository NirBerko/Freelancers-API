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
