package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sKush-1/magic_stream_movies_server/routes"
)

func main() {
	router := gin.Default()
	routes.SetupUnrotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
