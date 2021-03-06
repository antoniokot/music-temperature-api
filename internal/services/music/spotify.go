package music

import (
	"errors"
	"math/rand"
	"time"

	"github.com/antoniokot/music-temperature-api/config"
	"github.com/zmb3/spotify"
)

func GetPlaylistByTemperature(temp float32) (spotify.FullPlaylist, error) {

	var playlistID string

	if temp > 30 {
		partyPlaylists := []string {"37i9dQZF1DX8mBRYewE6or", "37i9dQZF1DX4MdXmAY6EDq", "37i9dQZF1DWWmaszSfZpom"}

		rand.Seed(time.Now().UnixNano())
		playlistID = partyPlaylists[rand.Intn(3)]

	} else if temp >= 15 { 
		popPlaylists := []string {"37i9dQZF1DX6aTaZa0K6VA", "37i9dQZF1DWVLcZxJO5zyf", "0akcrslPTLqkIRt6gfBKUi"}

		rand.Seed(time.Now().UnixNano())
		playlistID = popPlaylists[rand.Intn(3)]

	} else if temp < 15 && temp >= 10 {
		rockPlaylists := []string {"37i9dQZF1DWXRqgorJj26U", "37i9dQZF1DX4vCk1GJH7zl", "4CmVkOQJA12pmQwsICaYpC"}

		rand.Seed(time.Now().UnixNano())
		playlistID = rockPlaylists[rand.Intn(3)]

	} else if temp < 10 {
		classicPlaylists := []string {"37i9dQZF1DWWEJlAGA9gs0", "37i9dQZF1DWXmUP9DYdke2", "37i9dQZF1DWXzMRSBIsJEP"}

		rand.Seed(time.Now().UnixNano())
		playlistID = classicPlaylists[rand.Intn(3)]
	}

	id := spotify.ID(playlistID)
	playlist, err := config.SpotifyClient.GetPlaylist(id)

	if err != nil {
		return *playlist, errors.New("erro ao tentar recuperar a playlist")
	}

	return *playlist, nil
}