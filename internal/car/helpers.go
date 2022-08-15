package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	configs "ria-bot/configs"
)

func GetDataFromApi(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func ParseCarsResult(carsSlice []Car, cfg *configs.Config) []string {
	var carResult []string

	for _, car := range carsSlice {
		carInfo := fmt.Sprintf(
			"Title: %s\nMark: %s\nModel: %s\nYear: %d\nPrice: %d$\nLink: %s\nDescription: %s\n",
			car.Title, car.Mark, car.Model, car.CarData.Year, car.Price, cfg.RiaBaseUrl+car.Link, truncateText(car.CarData.Description, 200),
		)
		carResult = append(carResult, carInfo)
	}

	return carResult
}

func truncateText(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:strings.LastIndex(s[:max], " ")] + "..."
}
