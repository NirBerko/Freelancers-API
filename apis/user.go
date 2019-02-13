package apis

import (
	"freelancers/app"
	"freelancers/errors"
	"freelancers/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	userService interface {
		GetUserDetails(rs app.RequestScope) util.ResultParser
	}

	userResource struct {
		service userService
	}
)

func ServeUserResource(rg *gin.RouterGroup, service userService) {
	r := userResource{service}

	rg.GET("", r.GetUserDetails)
}

func (r *userResource) GetUserDetails(c *gin.Context) {
	result := r.service.GetUserDetails(app.GetRequestScope(c))

	if result.Error != nil {
		errorHandler := errors.InternalServerError(result.Error)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		c.JSON(http.StatusOK, result.Data)
	}
}
