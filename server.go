package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"os"
)

/*
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
}*/
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, heroku")
}
func main() {

	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
	/*
		app.LoadConfig()

		db, _ := gorm.Open("postgres", app.Config.DSN)
		AutoMigrate(db)

		r := gin.Default()
		gin.SetMode(app.Config.Mode)

		buildRouter(r, db)
		r.Run(":" + strconv.Itoa(app.Config.Server.Port))*/
}

/*
func buildRouter(router *gin.Engine, db *gorm.DB) {
	router.Use(
		app.Init(),
		app.Transactional(db),
	)

	router.GET("/ping", func(c *gin.Context)
		c.Abort()
		c.String(200, "OK "+app.Version)
	})

	authDao := daos.NewAuthDAO()ss
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

}*/
