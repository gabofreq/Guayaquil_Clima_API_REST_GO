package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
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
	Rain    float64
	Snow    float64
	Clouds  clouds_
	Coord   coord_
	Main    main_
	Wind    wind_
	Sys     sys_
	Weather []weather_
}

type main_ struct {
	Temp       float64
	Feels_like float64
	Temp_min   float64
	Temp_max   float64
	Pressure   float64
	Humidity   float64
	Sea_level  float64
	Grnd_level float64
}

type coord_ struct {
	Lat  float64
	Long float64
}

type clouds_ struct {
	All int
}

type wind_ struct {
	Speed float64
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
	Codigo             string
	Cuenta             int
	Id_Lista           int
	Pais               string
	Ciudad             string
	Latitud            float64
	Longitud           float64
	Dt                 int
	Temperatura        float64
	Sensacion_Termica  float64
	Temperatura_Minima float64
	Temperatura_Maxima float64
	Presion            float32
	Humedad            float32
	Nivel_Mar          float32
	Nivel_Suelo        float32
	Precipitacion      float32
	Nieve              float32
	Nubes              int
	Velocidad_Viento   float32
	Direccion_Viento   int
	Id_Clima           int
	Estado_Clima       string
	Descripcion_Clima  string
	Icono              string
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
	var variable = new(variables)
	variable.Codigo = t.Cod
	variable.Cuenta = t.Count
	variable.Id_Lista = (t.List[0]).Id
	variable.Pais = ((t.List[0]).Sys).Country
	variable.Ciudad = (t.List[0]).Name
	variable.Latitud = ((t.List[0]).Coord).Lat
	variable.Longitud = ((t.List[0]).Coord).Long
	variable.Dt = (t.List[0]).Dt

	variable.Temperatura = math.Round((((t.List[0]).Main).Temp-273.15)*100) / 100
	variable.Sensacion_Termica = math.Round((((t.List[0]).Main).Feels_like-273.15)*100) / 100
	variable.Temperatura_Minima = math.Round((((t.List[0]).Main).Temp_min-273.15)*100) / 100
	variable.Temperatura_Maxima = math.Round((((t.List[0]).Main).Temp_max-273.15)*100) / 100
	variable.Presion = ((t.List[0]).Main).Pressure
	variable.Humedad = ((t.List[0]).Main).Humidity
	variable.Nivel_Mar = ((t.List[0]).Main).Sea_level
	variable.Nivel_Suelo = ((t.List[0]).Main).Grnd_level
	variable.Precipitacion = (t.List[0]).Rain
	variable.Nieve = (t.List[0]).Snow
	variable.Nubes = ((t.List[0]).Clouds).All
	variable.Velocidad_Viento = ((t.List[0]).Wind).Speed
	variable.Direccion_Viento = ((t.List[0]).Wind).Deg
	variable.Id_Clima = ((t.List[0]).Weather[0]).Id
	variable.Estado_Clima = ((t.List[0]).Weather[0]).Main
	variable.Descripcion_Clima = ((t.List[0]).Weather[0]).Description
	variable.Icono = ((t.List[0]).Weather[0]).Icon
	jsons, err := json.Marshal(variable)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", string(jsons))
	json.Unmarshal(jsons, &variable)
}
