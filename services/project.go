package services

import (
	"errors"
	"freelancers/app"
	"freelancers/models"
	"freelancers/models/UIModels"
	"freelancers/util"
)

type (
	projectDao interface {
		CreateProject(rs app.RequestScope, project *models.Project) error
		GetProjectByID(rs app.RequestScope, id uint) models.Project
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

func (s *ProjectService) GetProjectByID(rs app.RequestScope, id uint) util.ResultParser {
	findProject := s.dao.GetProjectByID(rs, id)

	var skills []string

	for _, skill := range findProject.GetSkills() {
		skills = append(skills, skill.GetName())
	}

	project := UIModels.Project{
		ID:          findProject.GetID(),
		Title:       findProject.GetTitle(),
		Description: findProject.GetDescription(),
		BudgetType:  findProject.GetBudgetType(),
		BudgetLevel: findProject.GetBudgetLevel(),
		Skills:      skills,
	}

	var err error

	if findProject.GetID() == 0 {
		err = errors.New("Not Found")
	}

	return util.ResultParser{
		Data:   project,
		IsDone: true,
		Error:  err,
	}
}
