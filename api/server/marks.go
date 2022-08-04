package server

import (
	"encoding/json"
	"log"
)

type Mark struct {
	Name  string
	Value int64
}

func getAllMarks() []Mark {
	url := GetAllMarksAPI()
	responseData := GetDataFromApi(url)

	var m []Mark
	if err := json.Unmarshal(responseData, &m); err != nil {
		log.Fatal(err)
	}

	return m
}

func GetCarsByMark(mark string) []string {
	marks := getAllMarks()

	var searchableMark int64

	for i := range marks {
		if marks[i].Name == mark {
			searchableMark = marks[i].Value
		}
	}

	url := GetCarsByMarkAPI(searchableMark)
	responseData := GetDataFromApi(url)

	cars := GetCarsResult(responseData)

	return cars
}
