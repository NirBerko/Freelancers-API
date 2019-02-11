package apis

import (
	"freelancers/app"
	"freelancers/errors"
	"freelancers/models"
	"github.com/gin-gonic/gin"
)

type (
	projectService interface {
		CreateProject(rs app.RequestScope, project *models.Project) error
	}

	projectResource struct {
		service projectService
	}
)

func ServeProjectResource(rg *gin.RouterGroup, service projectService) {
	r := &projectResource{service}

	rg.PUT("/", r.CreateProject)
}

func (r *projectResource) CreateProject(c *gin.Context) {
	var project models.Project
	c.BindJSON(&project)

	err := r.service.CreateProject(app.GetRequestScope(c), &project)

	if err != nil {
		errorHandler := errors.InternalServerError(err)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(200, project)
	}
}
