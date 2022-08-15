package server

import (
	"encoding/json"
	"log"
	"sync"

	configs "ria-bot/configs"
)

type Result struct {
	SearchResult struct {
		CarIds struct {
			Ids []string `json:"ids"`
		} `json:"search_result"`
	} `json:"result"`
}

type Car struct {
	UserId  int64  `json:"userId"`
	City    string `json:"locationCityName"`
	Date    string `json:"addDate"`
	Price   int64  `json:"USD"`
	CarData struct {
		Description string `json:"description"`
		Year        int64  `json:"year"`
		Race        int64  `json:"raceInt"`
		Fuel        string `json:"fuelName"`
		IsSold      bool   `json:"isSold"`
	} `json:"autoData"`
	Mark  string `json:"markName"`
	Model string `json:"modelName"`
	Link  string `json:"linkToView"`
	Title string `json:"title"`
}

type Mark struct {
	Name  string
	Value int64
}

func GetAllCars(cfg *configs.Config) []string {
	url := GetAllCarsAPI(cfg)
	responseData := GetDataFromApi(url)

	cars := GetCarsResult(responseData, cfg)

	return cars
}

func GetCarById(id string, cars chan Car, wg *sync.WaitGroup, cfg *configs.Config) {
	url := GetCarByIdAPI(id, cfg)
	responseData := GetDataFromApi(url)

	var v Car
	if err := json.Unmarshal(responseData, &v); err != nil {
		log.Fatal(err)
	}
	cars <- v
}

func GetCarsResult(responseData []byte, cfg *configs.Config) []string {
	var wg sync.WaitGroup

	var r Result
	if err := json.Unmarshal(responseData, &r); err != nil {
		log.Fatal(err)
	}

	var carsSlice []Car
	carsChan := make(chan Car)

	for _, id := range r.SearchResult.CarIds.Ids {
		wg.Add(1)
		go GetCarById(id, carsChan, &wg, cfg)
		wg.Done()

		msg1 := <-carsChan
		carsSlice = append(carsSlice, msg1)
	}

	wg.Wait()

	cars := ParseCarsResult(carsSlice, cfg)

	return cars
}

func getAllMarks(cfg *configs.Config) []Mark {
	url := GetAllMarksAPI(cfg)
	responseData := GetDataFromApi(url)

	var m []Mark
	if err := json.Unmarshal(responseData, &m); err != nil {
		log.Fatal(err)
	}

	return m
}

func GetCarsByMark(mark string, cfg *configs.Config) []string {
	marks := getAllMarks(cfg)

	var searchableMark int64

	for i := range marks {
		if marks[i].Name == mark {
			searchableMark = marks[i].Value
		}
	}

	url := GetCarsByMarkAPI(searchableMark, cfg)
	responseData := GetDataFromApi(url)

	cars := GetCarsResult(responseData, cfg)

	return cars
}
