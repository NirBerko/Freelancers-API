package main

import (
	"freelancers/apis"
	"freelancers/app"
	"freelancers/daos"
	"freelancers/models"
	"freelancers/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func main() {
	app.LoadConfig()

	db, _ := gorm.Open("postgres", app.Config.DSN)
	AutoMigrate(db)

	r := gin.Default()
	gin.SetMode(app.Config.Mode)

	buildRouter(r, db)
	r.Run(":" + strconv.Itoa(app.Config.Server.Port))
}

func buildRouter(router *gin.Engine, db *gorm.DB) {
	router.Use(
		app.Init(),
		app.Transactional(db),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.Abort()
		c.String(200, "OK "+app.Version)
	})

	authDAO := daos.NewAuthDAO()
	apis.ServeAuthResource(router, services.NewAuthService(authDAO))

	router.Use(
		app.JwtMiddleware(),
	)

	router.GET("/pingAuth", func(c *gin.Context) {
		c.Abort()
		c.String(200, "OK "+app.Version)
	})

}
