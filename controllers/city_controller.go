package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antoniokot/music-temperature-api/models"
	"github.com/gin-gonic/gin"
)

func GetCity(con *gin.Context) {

	name := con.Param("name")

	client := &http.Client{}

 	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather?q=" + name + "&appid=b77e07f479efe92156376a8b07640ced", nil)

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

	con.IndentedJSON(http.StatusOK, getMusicByTemperature(c.Main.Temp))
  return
}