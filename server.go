package main

import (
	"example/wiki-server/configs"
	"example/wiki-server/controller"
	"example/wiki-server/middlewares"
	"example/wiki-server/routes"
	"example/wiki-server/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	// gin.SetMode(gin.ReleaseMode)
	setupLogOutput()
	server := gin.New()
	go configs.ConnectDB()

	server.Static("/css", "./template/css")

	server.LoadHTMLGlob("templates/*.html")

	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	server.Use(gin.Recovery(), middlewares.Logger()) // middlewares.BasicAuth(),
	// gindump.Dump(),

	server.POST("/login", controller.Login)
	server.GET("/logout", controller.Logout)
	auth := server.Group("/api")
	{
		auth.Use(middlewares.Authentication())

		auth.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(
				200,
				videoController.FindAll())
		})

		auth.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(
					http.StatusBadRequest,
					gin.H{
						"error": err.Error(),
					},
				)
			} else {
				ctx.JSON(
					http.StatusCreated,
					gin.H{
						"message": "Video Input is valid!!",
					},
				)
			}
		})

		auth.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK!!",
			})
		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	routes.UserRoute(server)

	server.Run(":8080")
}
