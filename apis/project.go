package apis

import (
	"freelancers/app"
	"freelancers/dataModel"
	"freelancers/errors"
	"freelancers/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	projectService interface {
		CreateProject(rs app.RequestScope, project *models.Project) error
		GetProjectByID(rs app.RequestScope, id uint) dataModel.ResultParser
	}

	projectResource struct {
		service projectService
	}
)

type bodyProject struct {
	models.Project
	Skills []string
}

func ServeProjectResource(rg *gin.RouterGroup, service projectService) {
	r := &projectResource{service}

	rg.PUT("", r.CreateProject)
	rg.GET("/:id", r.GetProjectByID)
}

func (r *projectResource) GetProjectByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	result := r.service.GetProjectByID(app.GetRequestScope(c), uint(id))

	if result.Error != nil {
		errorHandler := errors.InternalServerError(result.Error)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(http.StatusOK, result.Data)
	}
}

func (r *projectResource) CreateProject(c *gin.Context) {
	var bodyProject bodyProject
	c.BindJSON(&bodyProject)

	var skillsModels []models.Skill

	for _, skill := range bodyProject.Skills {
		skillsModels = append(skillsModels, models.Skill{Name: skill})
	}

	project := &models.Project{
		Title:       bodyProject.Title,
		Description: bodyProject.Description,
		BudgetLevel: bodyProject.BudgetLevel,
		BudgetType:  bodyProject.BudgetType,
		Skills:      skillsModels,
	}

	err := r.service.CreateProject(app.GetRequestScope(c), project)

	if err != nil {
		errorHandler := errors.InternalServerError(err)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(http.StatusOK, &project)
	}
}
