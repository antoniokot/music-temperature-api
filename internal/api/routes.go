package api

import (
	"github.com/gin-gonic/gin"
)

func StartRoutes() {
	router := gin.Default()
	
	router.GET("/playlist/city/:name", playlistRecomendation)

	router.Run(":3333")
}