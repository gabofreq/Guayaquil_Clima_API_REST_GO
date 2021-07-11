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
	Id    int
	Name  string
	Coord coord_
	Main  main_
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
	log.Printf("%d", t.Count)
	log.Println((t.List[0]).Name)
	log.Println(((t.List[0]).Main).Temp)
	log.Println(((t.List[0]).Coord).Lat)
	log.Println(((t.List[0]).Coord).Long)
}
