package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/antoniokot/music-temperature-api/config"
	"github.com/antoniokot/music-temperature-api/models"
	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)

func GetCity(con *gin.Context) {

	name := con.Param("name")

	temp, err := getTemperatureFromRedis(name)

	if err != nil || temp == -1 {

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

		temp = c.Main.Temp

		err = storeTemperatureInRedis(name, temp)

		if err != nil {
			boom.Internal(con.Writer, "Erro ao armazenar dados em cache")
			return
		}
	}

	playlist, err := getPlaylistByTemperature(temp) 

	if err != nil {
		boom.Internal(con.Writer, err)
		return
	}

	con.IndentedJSON(http.StatusOK, playlist)
}

func getTemperatureFromRedis(name string) (float32, error) {

	val, err := config.RedisClient.Get(name).Result()

	if err != nil {
		return -1, err
	}

	fVal, err := strconv.ParseFloat(val, 32)

	if err != nil {
		return -1, err
	}

	return float32(fVal), nil
}

func storeTemperatureInRedis(key string, val float32) error {

	err := config.RedisClient.Set(key, val, 300000000000).Err()

	return err
}