package apis

import (
	"freelancers/app"
	"freelancers/models/UIModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	userService interface {
		GetUserDetails(rs app.RequestScope) UIModels.User
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
	user := r.service.GetUserDetails(app.GetRequestScope(c))

	c.JSON(http.StatusOK, user)
}
