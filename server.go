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
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := gorm.Open("postgres", os.Getenv("DSN"))
	AutoMigrate(db)

	r := gin.Default()
	gin.SetMode(os.Getenv("MODE"))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

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
		c.String(http.StatusOK, "OK "+app.Version)
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
