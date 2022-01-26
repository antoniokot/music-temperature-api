package controllers

import (
	"testing"
	"fmt"
)

func TestGetMusicByTemperatureParty(t *testing.T)  {
	pl, err := getPlaylistByTemperature(31)

	if err == nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de PARTY, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperaturePop(t *testing.T)  {
	pl, err := getPlaylistByTemperature(30)

	if err == nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de POP, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperatureRock(t *testing.T)  {
	pl, err := getPlaylistByTemperature(10)

	if err == nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de ROCK, %v, error`, err)
	}

	fmt.Println(pl.Name)
}

func TestGetMusicByTemperatureClassic(t *testing.T)  {
	pl, err := getPlaylistByTemperature(9)

	if err == nil {
		t.Fatalf(`Não foi possível recuperar uma playlist de CLASSICA, %v, error`, err)
	}

	fmt.Println(pl.Name)
}