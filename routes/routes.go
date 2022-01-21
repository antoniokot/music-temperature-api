package routes

import (
	"github.com/gin-gonic/gin"
)

func startRoutes() {
	router := gin.Default()
	
	router.GET("/music/city/:name", cityController.GetCity)

	router.Run("localhost:3333")
}