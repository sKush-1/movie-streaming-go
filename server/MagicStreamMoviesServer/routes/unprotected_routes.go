package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sKush-1/magic_stream_movies_server/controllers"
)

func SetupUnrotectedRoutes(router *gin.Engine) {
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.GET("/movies", controller.GetMovies())

}
