package services

import (
	"freelancers/app"
	"freelancers/models"
)

type (
	projectDao interface {
		CreateProject(rs app.RequestScope, project *models.Project) error
	}

	ProjectService struct {
		dao projectDao
	}
)

func NewProjectService(dao projectDao) *ProjectService {
	return &ProjectService{dao}
}

func (s *ProjectService) CreateProject(rs app.RequestScope, project *models.Project) error {
	return s.dao.CreateProject(rs, project)
}
