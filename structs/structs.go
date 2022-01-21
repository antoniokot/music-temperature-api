package structs

type coord struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type weather struct {
	ID int16 `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type main struct {	
	Temp int16 `json:"temp"`
	Temp_min int16 `json:"temp_min"`
	Temp_max int16 `json:"temp_max"`
	Feels_like int16 `json:"feels_like"`
	Pressure int16 `json:"pressure"`
	Humidity int16 `json:"humidity"`
	Sea_level int16 `json:"sea_level"`
	Grnd_level int16 `json:"grnd_level"`
}

type wind struct {
	Speed float32 `json:"speed"`
	Deg int32 `json:"deg"`
	Gust float32 `json:"gust"`
}

type clouds struct {
	All int8 `json:"all"`
}

type sys struct {
	ID int32 `json:"id"`
	Type int8 `json:"type"`
	country string `json:"country"`
	sunrise int32 `json:"sunrise"`
	sunset int32 `json:"sunset"`		
}

type city struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
	Cod int32 `json:"cod"`
	Timezone int32 `json:"timezone"`

	Sys sys `json:"sys"`

	Dt int32 `json:"dt"`

	Wind wind `json:"wind"`

	Visibility int16 `json:"visibility"`

	Main main `json:"main"`

	Base string `json:"base"`

	Weather []weather `json:"weather"`

	Coord coord `json:"coord"`
}

type music struct {

}