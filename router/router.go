package router

import (
	"{{projectName}}/app"
	_ "{{projectName}}/docs"
	"{{projectName}}/migrate"
	"{{projectName}}/controller"
	"{{projectName}}/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Start() {
	app.InitConfig()
	app.InitDB()
	migrate.CreateTable()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//!!do not delete gen will generate router code at here

	addr := app.Config.Http.Domain + ":" + app.Config.Http.Port

	r.Run(addr) // listen and serve on 0.0.0.0:8080
}
