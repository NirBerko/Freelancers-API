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
