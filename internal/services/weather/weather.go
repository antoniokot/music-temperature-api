package weather

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/antoniokot/music-temperature-api/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetWeatherByCityName(name string, con *gin.Context) (float32, error) {

	// temp, err := models.GetTemperatureFromRedis(name)

	// if err != nil || temp == -1 {

		client := &http.Client{}

		req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather?q=" + name + "&appid=b77e07f479efe92156376a8b07640ced", nil)

		if err != nil {
			return -1, errors.New("Erro ao montar requisição")
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)

		if resp.StatusCode != 200 || err != nil {
			return -1, errors.New("Erro ao recuperar cidade. O nome da cidade "+name+" está certo?")
		}
		
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		
		if err != nil {
			return -1, errors.New("Erro ao recuperar resposta do serviço")
		}

		var c models.City
		json.Unmarshal(bodyBytes, &c)

		temp := c.Main.Temp

		// temp = c.Main.Temp

		// err = models.StoreTemperatureInRedis(name, temp)

		if err != nil {
			return -1, errors.New("Erro ao armazenar dados em cache")
		}
	// }

	return float32(math.Round(float64(temp - 273.15) * 100) / 100), nil
}