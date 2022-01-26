package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"

	"github.com/antoniokot/music-temperature-api/models"
	"github.com/gin-gonic/gin"
)

func GetCity(con *gin.Context) {

	name := con.Param("name")

	client := &http.Client{}

 	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather?q=" + name + "&appid=b77e07f479efe92156376a8b07640ced", nil)

 	if err != nil {
		log.Fatal("Erro ao montar requisição")
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Erro ao recuperar tempo da cidade")
	}
	
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal("Erro ao recuperar resposta do serviço")
	}

	var c models.City
	json.Unmarshal(bodyBytes, &c)

	playlist, err := getPlaylistByTemperature(c.Main.Temp) 

	if err != nil {
		log.Fatal(err)
	}

	con.IndentedJSON(http.StatusOK, playlist)
}