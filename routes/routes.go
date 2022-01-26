package routes

import (
	"github.com/antoniokot/music-temperature-api/controllers"
	"github.com/gin-gonic/gin"
)

func StartRoutes() {
	router := gin.Default()
	
	router.GET("/playlist/city/:name", controllers.GetCity)

	router.Run("localhost:3333")
}