package main

import (
	"fmt"

	"github.com/antoniokot/music-temperature-api/config"
	"github.com/antoniokot/music-temperature-api/internal/api"
)

func main() {
	
	fmt.Println("Subindo servidor em http://localhost:3333")

	config.StartRedisConfig()
	config.StartSpotifyConfig()
	api.StartRoutes()
}