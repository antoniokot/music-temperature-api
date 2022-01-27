package controllers

import (
	"fmt"
	"testing"

	"github.com/antoniokot/music-temperature-api/config"
)

func TestGetMusicByTemperatureParty(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := getPlaylistByTemperature(310)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de PARTY, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperaturePop(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := getPlaylistByTemperature(300)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de POP, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperatureRock(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := getPlaylistByTemperature(284)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de ROCK, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperatureClassic(t *testing.T)  {

	config.StartSpotifyConfig()

	pl, err := getPlaylistByTemperature(9)

	if err != nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de CLASSICA, %v, error`, err)
	}

	fmt.Println(pl.Name)
}