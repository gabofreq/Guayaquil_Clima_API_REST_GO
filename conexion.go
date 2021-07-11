package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	api_get_information()
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
