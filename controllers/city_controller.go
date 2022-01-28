package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/antoniokot/music-temperature-api/models"
	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)

func GetCity(con *gin.Context) {

	name := con.Param("name")

	

	client := &http.Client{}

 	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather?q=" + name + "&appid=b77e07f479efe92156376a8b07640ced", nil)

 	if err != nil {
		boom.Internal(con.Writer, "Erro ao montar requisição")
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if resp.StatusCode != 200 || err != nil {
		boom.NotFound(con.Writer, "Erro ao recuperar cidade. O nome da cidade "+name+" está certo?")
		return
	}
	
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		boom.Internal(con.Writer, "Erro ao recuperar resposta do serviço")
		return
	}

	var c models.City
	json.Unmarshal(bodyBytes, &c)

	playlist, err := getPlaylistByTemperature(c.Main.Temp) 

	if err != nil {
		boom.Internal(con.Writer, err)
		return
	}

	con.IndentedJSON(http.StatusOK, playlist)
}