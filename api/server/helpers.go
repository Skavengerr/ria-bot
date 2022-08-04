package server

import (
	"fmt"
	"strings"

	"ria-bot/configs"
)

func ParseCarsResult(carsSlice []Car) []string {
	var carResult []string

	for _, car := range carsSlice {
		carInfo := fmt.Sprintf(
			"Title: %s\nMark: %s\nModel: %s\nYear: %d\nPrice: %d$\nLink: %s\nDescription: %s\n",
			car.Title, car.Mark, car.Model, car.Car.Year, car.Price, configs.RIA_BASE_URL+car.Link, truncateText(car.Car.Description, 200),
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
