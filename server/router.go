package server

import (
	"github.com/Arturo0911/go-android/controllers"
	. "github.com/Arturo0911/go-android/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {

	repoLayer := controllers.New()
	router.Group("/realtime").
		GET("/get-users", repoLayer.GetUsersControllers).
		POST("/post-users", repoLayer.PostUsersControllers)

	router.Use(func(c *gin.Context) {
		c.JSON(404, JsonResponse{
			Success: false,
			Error:   "resource not found!!",
		})
	})

	return router
}
