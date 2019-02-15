package apis

import (
	"freelancers/app"
	"freelancers/errors"
	"freelancers/models"
	"freelancers/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	projectService interface {
		CreateProject(rs app.RequestScope, project *models.Project) util.HandleResult
		GetProjectByID(rs app.RequestScope, id uint) util.HandleResult
		GetAllProjects(rs app.RequestScope) util.HandleResult
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
	rg.GET("/getbyid/:id", r.GetProjectByID)
	rg.GET("/all", r.GetAllProjects)
}

func (r *projectResource) GetProjectByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	result := r.service.GetProjectByID(app.GetRequestScope(c), uint(id))

	if result.Error != nil {
		errorHandler := errors.InternalServerError(result.Error)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(http.StatusOK, result.GetResult())
	}
}

func (r *projectResource) CreateProject(c *gin.Context) {
	var bodyProject bodyProject
	c.BindJSON(&bodyProject)

	var skillsModels []models.Skill

	for _, skill := range bodyProject.Skills {
		skillsModels = append(skillsModels, models.Skill{Name: skill})
	}

	rs := app.GetRequestScope(c)

	project := &models.Project{
		Title:       bodyProject.Title,
		Description: bodyProject.Description,
		BudgetLevel: bodyProject.BudgetLevel,
		BudgetType:  bodyProject.BudgetType,
		Skills:      skillsModels,
		User:        rs.User(),
	}

	result := r.service.CreateProject(rs, project)

	if result.Error != nil {
		errorHandler := errors.InternalServerError(result.Error)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(http.StatusOK, result.GetResult())
	}
}

func (r *projectResource) GetAllProjects(c *gin.Context) {
	result := r.service.GetAllProjects(app.GetRequestScope(c))

	if result.GetError() != nil || !result.GetIsSuccess() {
		errorHandler := errors.InternalServerError(result.Error)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(200, result.GetResult())
	}
}
