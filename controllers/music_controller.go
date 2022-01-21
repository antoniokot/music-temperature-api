package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antoniokot/music-temperature-api/models"
)

func getMusicByTemperature(temp float32) models.Music {
	
	client := &http.Client{}

 	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

 	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}
	
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Print(err.Error())
	}

	var c models.City
	json.Unmarshal(bodyBytes, &c)

	var m models.Music

	return m
}