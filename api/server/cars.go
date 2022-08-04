package server

import (
	"encoding/json"
	"log"
	"sync"
)

type Result struct {
	Result SearchResult `json:"result"`
}

type SearchResult struct {
	SearchResult CarIds `json:"search_result"`
}

type CarIds struct {
	Ids []string `json:"ids"`
}

type Car struct {
	UserId int64   `json:"userId"`
	City   string  `json:"locationCityName"`
	Date   string  `json:"addDate"`
	Price  int64   `json:"USD"`
	Car    CarData `json:"autoData"`
	Mark   string  `json:"markName"`
	Model  string  `json:"modelName"`
	Link   string  `json:"linkToView"`
	Title  string  `json:"title"`
}

type CarData struct {
	Description string `json:"description"`
	Year        int64  `json:"year"`
	Race        int64  `json:"raceInt"`
	Fuel        string `json:"fuelName"`
	IsSold      bool   `json:"isSold"`
}

func GetAllCars() []string {
	url := GetAllCarsAPI()
	responseData := GetDataFromApi(url)

	cars := GetCarsResult(responseData)

	return cars
}

func GetCarById(id string, cars chan Car, wg *sync.WaitGroup) {
	url := GetCarByIdAPI(id)
	responseData := GetDataFromApi(url)

	var v Car
	if err := json.Unmarshal(responseData, &v); err != nil {
		log.Fatal(err)
	}
	cars <- v
}

func GetCarsResult(responseData []byte) []string {
	var wg sync.WaitGroup

	var r Result
	if err := json.Unmarshal(responseData, &r); err != nil {
		log.Fatal(err)
	}

	var carsSlice []Car
	carsChan := make(chan Car)

	for _, id := range r.Result.SearchResult.Ids {
		wg.Add(1)
		go GetCarById(id, carsChan, &wg)
		wg.Done()

		msg1 := <-carsChan
		carsSlice = append(carsSlice, msg1)
	}

	wg.Wait()

	cars := ParseCarsResult(carsSlice)

	return cars
}
