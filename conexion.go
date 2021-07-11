package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type nivel_0 struct {
	Message string
	Cod     string
	Count   int
	List    []nivel_1
}

type nivel_1 struct {
	Id      int
	Name    string
	Dt      int
	Rain    float32
	Snow    float32
	Clouds  clouds_
	Coord   coord_
	Main    main_
	Wind    wind_
	Sys     sys_
	Weather []weather_
}

type main_ struct {
	Temp       float32
	Feels_like float32
	Temp_min   float32
	Temp_max   float32
	Pressure   float32
	Humidity   float32
	Sea_level  float32
	Grnd_level float32
}

type coord_ struct {
	Lat  float32
	Long float32
}

type clouds_ struct {
	All int
}

type wind_ struct {
	Speed float32
	Deg   int
}

type sys_ struct {
	Country string
}

type weather_ struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type variables struct {
	Codigo              string
	Cuenta              int
	Id_Lista            int
	Ciudad              string
	Pais                string
	Latitud             float32
	Longitud            float32
	Temperatura         float32
	Sensaacion_Termica  float32
	Temperatura_Mininna float32
	Temperatura_Maxima  float32
	Presion             float32
	Humedad             float32
	Nivel_Mar           float32
	Nivel_Suelo         float32
	Precipitacion       float32
	Nieve               float32
	Nubes               int
	Velocidad_Viento    float32
	Direccion_Viento    int
	Clima_Id            int
	Estado_Clima        string
	Descripcion_Clima   string
	Icono               string
}

func main() {
	//api_get_information()
	api_get_clima()
}

func api_get_information() {
	url := "https://world-geo-data.p.rapidapi.com/countries/EC?language=en%2Cru%2Czh%2Ces%2Car%2Cfr%2Cfa%2Cja%2Cpl%2Cit%2Cpt%2Cde"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", "e80cbef646mshdc00b2f552d3ab5p128cf5jsnaf676b00aaae")
	req.Header.Add("x-rapidapi-host", "world-geo-data.p.rapidapi.com")
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(string(body))
}

func api_get_clima() {
	url := "https://openweathermap.org/data/2.5/find?q=guayaquil&appid=439d4b804bc8187953eb36d2a8c26a02&units=metric"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var t nivel_0
	json.Unmarshal(body, &t)
	log.Println(t.Message)
	log.Println(t.Cod)
	log.Println(t.Count)
	log.Println((t.List[0]).Id)
	log.Println((t.List[0]).Name)
	log.Println((t.List[0]).Dt)
	log.Println(((t.List[0]).Coord).Lat)
	log.Println(((t.List[0]).Coord).Long)
	log.Println(((t.List[0]).Main).Temp)
	log.Println(((t.List[0]).Main).Feels_like)
	log.Println(((t.List[0]).Main).Temp_min)
	log.Println(((t.List[0]).Main).Temp_max)
	log.Println(((t.List[0]).Main).Pressure)
	log.Println(((t.List[0]).Main).Humidity)
	log.Println(((t.List[0]).Main).Sea_level)
	log.Println(((t.List[0]).Main).Grnd_level)
	log.Println(((t.List[0]).Wind).Speed)
	log.Println(((t.List[0]).Wind).Deg)
	log.Println((t.List[0]).Rain)
	log.Println((t.List[0]).Snow)
	log.Println(((t.List[0]).Clouds).All)
	log.Println(((t.List[0]).Weather[0]).Id)
	log.Println(((t.List[0]).Weather[0]).Main)
	log.Println(((t.List[0]).Weather[0]).Description)
	log.Println(((t.List[0]).Weather[0]).Icon)
}
