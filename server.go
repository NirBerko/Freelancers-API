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
	db.AutoMigrate(&models.Skill{})
	db.Model(&models.Project{}).Related(&models.Skill{}, "Skills")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("DSN: " + os.Getenv("DSN"))

	db, _ := gorm.Open("postgres", os.Getenv("DSN"))

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	AutoMigrate(db)

	r := gin.Default()
	gin.SetMode(os.Getenv("MODE"))

	buildRouter(r, db)
	r.Run(":" + os.Getenv("PORT"))
}

func buildRouter(router *gin.Engine, db *gorm.DB) {
	router.Use(
		app.Init(),
		app.Transactional(db),
		CORSMiddleware(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.Abort()
		c.String(http.StatusOK, "OK "+app.Version)
	})

	authDao := daos.NewAuthDAO()
	apis.ServeAuthResource(router, services.NewAuthService(authDao))

	router.Use(
		app.JwtMiddleware(),
	)

	userDao := daos.NewUserDAO()
	apis.ServeUserResource(router.Group("/user"), services.NewUserService(userDao))

	projectDao := daos.NewProjectDAO()
	apis.ServeProjectResource(router.Group("/project"), services.NewProjectService(projectDao))

}
