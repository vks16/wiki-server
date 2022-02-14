package routes

import (
	"example/wiki-server/controller"
	"example/wiki-server/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	// All routes related to users comes here
	users := router.Group("/user")
	users.Use(middlewares.Authentication())
	users.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK, gin.H{
				"message": "User routes working",
			},
		)
	})

	users.POST("/all", controller.FindAll())
	users.GET("/:userId", controller.GetAUser())
	users.POST("/create", controller.CreateUser())
	users.PUT("/:userId", controller.EditAUser())
	users.DELETE("/:userId", controller.DeleteAUser())
}
