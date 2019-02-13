package apis

import (
	"freelancers/app"
	"freelancers/errors"
	"freelancers/models"
	"freelancers/models/UIModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	authService interface {
		Login(rs app.RequestScope, user *models.User) *UIModels.User
		Register(rs app.RequestScope, user *models.User) error
	}

	authResource struct {
		service authService
	}

	userWithToken struct {
		UIModels.User
		Token string `json:"token"`
	}
)

func ServeAuthResource(rg *gin.Engine, service authService) {
	r := &authResource{service}

	rg.POST("/login", r.Login)
	rg.POST("/register", r.Register)
}

func (r *authResource) Login(c *gin.Context) {
	var user models.User
	_ = c.BindJSON(&user)

	authenticatedUser := r.service.Login(app.GetRequestScope(c), &user)

	if authenticatedUser != nil {
		token, err := app.EasyNewJWT(user.GetID())

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.JSON(http.StatusOK, userWithToken{
			*authenticatedUser,
			token,
		})
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func (r *authResource) Register(c *gin.Context) {
	var user models.User
	_ = c.BindJSON(&user)

	err := r.service.Register(app.GetRequestScope(c), &user)

	if err != nil {
		errorHandler := errors.InternalServerError(err)
		c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
	} else {
		token, err := app.EasyNewJWT(user.GetID())

		if err != nil {
			errorHandler := errors.InternalServerError(err)
			c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Error())
		}

		c.JSON(http.StatusOK, userWithToken{UIModels.User{
			ID:        user.GetID(),
			Email:     user.GetEmail(),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
		}, token})
	}
}
