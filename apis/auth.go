package apis

import (
	"fmt"
	"freelancers/app"
	"freelancers/errors"
	"freelancers/models"
	"github.com/gin-gonic/gin"
)

type (
	authService interface {
		Login(rs app.RequestScope, user *models.User)
		Register(rs app.RequestScope, user *models.User) error
	}

	authResource struct {
		service authService
	}

	userWithToken struct {
		models.User
		Token string `json:"token"`
	}
)

func ServeAuthResource(rg *gin.Engine, service authService) {
	r := &authResource{service}

	rg.POST("/login", r.Auth)
	rg.POST("/register", r.Register)
}

func (r *authResource) Auth(c *gin.Context) {
	var user models.User
	_ = c.BindJSON(&user)

	//r.service.Login(app.GetRequestScope(c), &user)*/

	token, err := app.EasyNewJWT(user.GetID())

	if err != nil {
		fmt.Println("UNAUTHORIZED")
	}

	c.JSON(200, userWithToken{user, token})
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

		c.JSON(200, userWithToken{user, token})
	}
}
