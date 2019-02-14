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

func (s *ProjectService) CreateProject(rs app.RequestScope, project *models.Project) util.ResultParser {
	err := s.dao.CreateProject(rs, project)

	var skills []string

	for _, skill := range project.Skills {
		skills = append(skills, skill.GetName())
	}

	projectUI := UIModels.Project{
		ID:          project.GetID(),
		Title:       project.GetTitle(),
		Description: project.GetDescription(),
		BudgetLevel: project.GetBudgetLevel(),
		BudgetType:  project.GetBudgetType(),
		Skills:      skills,
		User: UIModels.User{
			ID:        project.GetUser().GetID(),
			Email:     project.GetUser().GetEmail(),
			FirstName: project.GetUser().GetFirstName(),
			LastName:  project.GetUser().GetLastName(),
		},
	}

	return util.ResultParser{
		Data:   projectUI,
		IsDone: true,
		Error:  err,
	}
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
		User: UIModels.User{
			ID:        findProject.GetUser().GetID(),
			Email:     findProject.GetUser().GetEmail(),
			FirstName: findProject.GetUser().GetFirstName(),
			LastName:  findProject.GetUser().GetLastName(),
		},
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
