package api

import (
	"github.com/antoniokot/music-temperature-api/internal/services/music"
	"github.com/antoniokot/music-temperature-api/internal/services/weather"
	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)

func playlistRecomendation(c *gin.Context) {
	
	name := c.Param("name")

	temp, err := weather.GetWeatherByCityName(name, c)

	if err != nil {
		boom.BadRequest(c.Writer, err)
		return
	}

	playlist, err := music.GetPlaylistByTemperature(temp)

	if err != nil {
		boom.BadRequest(c.Writer, err)
		return
	}

	c.JSON(200, gin.H{
		"temperature": temp,
		"name":        playlist.Name,
		"owner":			 playlist.Owner.DisplayName,
		"description": playlist.Description,
		"tracks": 		 playlist.Tracks.Total,
	})
}