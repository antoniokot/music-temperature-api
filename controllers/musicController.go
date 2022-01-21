package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMusicByTemperature(temp float32) music {
	
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

	var c city
	json.Unmarshal(bodyBytes, &c)

	var m music

	return m
}