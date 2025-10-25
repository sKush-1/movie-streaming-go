package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sKush-1/magic_stream_movies_server/controllers"
	"github.com/sKush-1/magic_stream_movies_server/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

}
