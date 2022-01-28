package main

import (
	"fmt"

	"github.com/antoniokot/music-temperature-api/config"
	"github.com/antoniokot/music-temperature-api/routes"
)

func main() {
	fmt.Println("Subindo servidor em http://localhost:3333")

	db := config.StartRedisConfig()
	x, err := db.Get("Curitiba").Result()

	fmt.Println(x, err)

	config.StartSpotifyConfig()
	routes.StartRoutes()
}