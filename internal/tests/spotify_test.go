package controllers

import (
	"fmt"
	"testing"

	"github.com/antoniokot/music-temperature-api/config"
	"github.com/antoniokot/music-temperature-api/internal/services/music"
)

func TestGetMusicByTemperatureParty(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := music.GetPlaylistByTemperature(310)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de PARTY, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperaturePop(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := music.GetPlaylistByTemperature(300)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de POP, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperatureRock(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := music.GetPlaylistByTemperature(284)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de ROCK, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperatureClassic(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := music.GetPlaylistByTemperature(9)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de CLASSICA, %v, error`, err)
	}

	fmt.Println(pl.Name)
}