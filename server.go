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
	"os"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
}

func main() {
	app.LoadConfig()

	db, _ := gorm.Open("postgres", app.Config.DSN)
	AutoMigrate(db)

	r := gin.Default()
	gin.SetMode(app.Config.Mode)

	buildRouter(r, db)
	r.Run(":" + os.Getenv("PORT"))
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

	authDao := daos.NewAuthDAO()
	apis.ServeAuthResource(router, services.NewAuthService(authDao))

	projectDAO := daos.NewProjectDAO()
	apis.ServeProjectResource(router.Group("/project"), services.NewProjectService(projectDAO))

	router.Use(
		app.JwtMiddleware(),
	)

	router.GET("/pingAuth", func(c *gin.Context) {
		c.Abort()
		c.String(200, "OK "+app.Version)
	})

}
