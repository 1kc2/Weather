package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getCurrentWeather(location string, token string) {
	// creates a http client, its job is to transmit and receive HTTP messages
	client := &http.Client{}

	// NewRequest builds an HTTP request and returns it
	// Takes a method, URL
	req, err := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/weather", nil)

	// takes the above new request and adds queries to the end of it
	q := req.URL.Query()
	q.Add("q", location)
	q.Add("units", "imperial")
	q.Add("appid", token)
	// RawQuery encodes query values without '?'
	// Encode creates the request sring
	req.URL.RawQuery = q.Encode()

	// takes a request and executes it using the http client
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// If Response.Body won't be closed with Close() method than a resources 	associated with a fd won't be freed. This is a resource leak.
	defer resp.Body.Close()
	// reads from r until an error or EOF and returns the data it read
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// parses the json from body and stores it in the pointer
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Printf("The current weather in 	%s:\n-------------------------------------\nCurrent Temp: %.1f\nToday's Low: %.1f\nToday's High: %.1f\nHumidity: %d\nDescription: %v\n", currentWeather.Name, currentWeather.Main.Temp, currentWeather.Main.TempMin, currentWeather.Main.TempMax, currentWeather.Main.Humidity, currentWeather.Weather[1])

}

func getForecast(location string, token string) {
	// creates a http client, its job is to transmit and receive HTTP messages
	client := &http.Client{}

	// NewRequest builds an HTTP request and returns it
	// Takes a method, URL
	req, err := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/forecast", nil)

	// takes the above new request and adds queries to the end of it
	q := req.URL.Query()
	q.Add("q", location)
	q.Add("units", "imperial")
	q.Add("appid", token)
	// RawQuery encodes query values without '?'
	// Encode creates the request sring
	req.URL.RawQuery = q.Encode()

	// takes a request and executes it using the http client
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// If Response.Body won't be closed with Close() method than a resources 	associated with a fd won't be freed. This is a resource leak.
	defer resp.Body.Close()
	// reads from r until an error or EOF and returns the data it read
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// parses the json from body and stores it in the pointer
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Printf("The forecase is for %s is:\n-------------------------------------\nDate: %s\nForecast: %v\n", forecast.List.dt, forecast.List.Main)
}

// Struct to hold the values of the API reponse
var currentWeather = struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}{}

var forecast = struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		} `json:"wind"`
		Rain struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
	} `json:"city"`
}{}

func main() {
	// make sure a location was given
	//if len(os.Args) != 2 {
	//	log.Print("Please provide a location")
	//		os.Exit(1)
	//	}
	location := os.Args[1]
	token := "8f6bed5d8fb82eab0589d3b119cec424"

}
