package controllers

import (
	"fmt"
	"log"

	"github.com/antoniokot/music-temperature-api/config"
	"github.com/zmb3/spotify"
)

func getPlaylistByID(id string) spotify.FullPlaylist {

	playlistID := spotify.ID(id)
	playlist, err := config.Client.GetPlaylist(playlistID)

	fmt.Println(*playlist)

	if err != nil {
		log.Fatalf("Erro ao tentar recuperar a playlist: %v", err)
	}

	return *playlist
}