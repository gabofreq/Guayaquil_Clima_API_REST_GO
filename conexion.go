package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type nivel_clim_0 struct {
	Message string
	Cod     string
	Count   int
	List    []clima
}

type nivel_info_0 struct {
	Geonameid  int
	Name       string
	Population int
	Latitude   float64
	Longitude  float64
	Wiki_id    string
	Wiki_url   string
	Division   division_
	Country    country_
	Currency   currency_
	Timezone   timezone_
}

type division_ struct {
	Code      string
	Geonameid int
	Name      string
	Type      string
}

type country_ struct {
	Code       string
	Geonameid  int
	Name       string
	Depends_on string
}

type currency_ struct {
	Code string
	Name string
}

type timezone_ struct {
	Timezone   string
	Time       string
	Gtm_offset int
}

type clima struct {
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
	Geonameid_Pais     int
	Codigo_Pais        string
	Pais_Nombre        string
	Geonameid_Ciudad   int
	Ciudad             string
	Geonameid_Canton   int
	Canton             string
	Poblacion          int
	Wiki_id            string
	Wiki_url           string
	Code_Division      string
	Codigo_Moneda      string
	Moneda             string
	Zona_Horaria       string
	Fecha              string
	Gtm_offset         int
	Latitud            float64
	Longitud           float64
	Dt                 int
	Temperatura        float64
	Sensacion_Termica  float64
	Temperatura_Minima float64
	Temperatura_Maxima float64
	Presion            float64
	Humedad            float64
	Nivel_Mar          float64
	Nivel_Suelo        float64
	Precipitacion      float64
	Nieve              float64
	Nubes              int
	Velocidad_Viento   float64
	Direccion_Viento   int
	Id_Clima           int
	Estado_Clima       string
	Descripcion_Clima  string
	Icono              string
}

func main() {
	api_get_info()
}

func api_get_info() {
	url_1 := "https://openweathermap.org/data/2.5/find?q=guayaquil&appid=439d4b804bc8187953eb36d2a8c26a02&units=metric"
	req_1, _ := http.NewRequest("GET", url_1, nil)
	res_1, _ := http.DefaultClient.Do(req_1)
	body_1, _ := ioutil.ReadAll(res_1.Body)
	defer res_1.Body.Close()
	var t1 nivel_clim_0
	json.Unmarshal(body_1, &t1)
	url_2 := "https://world-geo-data.p.rapidapi.com/cities/3657509"
	req_2, _ := http.NewRequest("GET", url_2, nil)
	req_2.Header.Add("x-rapidapi-key", "d1fbafd1a0msh396241fdf61ab74p1d4c23jsn8734b01c4b06")
	req_2.Header.Add("x-rapidapi-host", "world-geo-data.p.rapidapi.com")
	res_2, _ := http.DefaultClient.Do(req_2)
	body_2, _ := ioutil.ReadAll(res_2.Body)
	defer res_2.Body.Close()
	var t2 nivel_info_0
	json.Unmarshal(body_2, &t2)
	var variable = new(variables)
	variable.Codigo = t1.Cod
	variable.Cuenta = t1.Count
	variable.Id_Lista = (t1.List[0]).Id
	variable.Codigo_Pais = (t2.Country).Code
	variable.Geonameid_Pais = (t2.Country).Geonameid
	variable.Pais_Nombre = (t2.Country).Name
	variable.Geonameid_Ciudad = t2.Geonameid
	variable.Ciudad = t2.Name
	variable.Geonameid_Canton = (t2.Division).Geonameid
	variable.Canton = (t2.Division).Name
	variable.Latitud = t2.Latitude
	variable.Longitud = t2.Longitude
	variable.Poblacion = t2.Population
	variable.Wiki_id = t2.Wiki_id
	variable.Wiki_url = t2.Wiki_url
	variable.Codigo_Moneda = (t2.Currency).Code
	variable.Moneda = (t2.Currency).Name
	variable.Zona_Horaria = (t2.Timezone).Timezone
	variable.Fecha = (t2.Timezone).Time[:19]
	variable.Gtm_offset = (t2.Timezone).Gtm_offset
	variable.Dt = (t1.List[0]).Dt
	variable.Temperatura = math.Round((((t1.List[0]).Main).Temp-273.15)*100) / 100
	variable.Sensacion_Termica = math.Round((((t1.List[0]).Main).Feels_like-273.15)*100) / 100
	variable.Temperatura_Minima = math.Round((((t1.List[0]).Main).Temp_min-273.15)*100) / 100
	variable.Temperatura_Maxima = math.Round((((t1.List[0]).Main).Temp_max-273.15)*100) / 100
	variable.Presion = ((t1.List[0]).Main).Pressure
	variable.Humedad = ((t1.List[0]).Main).Humidity
	variable.Nivel_Mar = ((t1.List[0]).Main).Sea_level
	variable.Nivel_Suelo = ((t1.List[0]).Main).Grnd_level
	variable.Precipitacion = (t1.List[0]).Rain
	variable.Nieve = (t1.List[0]).Snow
	variable.Nubes = ((t1.List[0]).Clouds).All
	variable.Velocidad_Viento = ((t1.List[0]).Wind).Speed
	variable.Direccion_Viento = ((t1.List[0]).Wind).Deg
	variable.Id_Clima = ((t1.List[0]).Weather[0]).Id
	variable.Estado_Clima = ((t1.List[0]).Weather[0]).Main
	variable.Descripcion_Clima = ((t1.List[0]).Weather[0]).Description
	variable.Icono = ((t1.List[0]).Weather[0]).Icon
	jsons, err := json.Marshal(variable)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", string(jsons))
	json.Unmarshal(jsons, &variable)
}
