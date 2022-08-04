package server

import (
	"fmt"
	"ria-bot/configs"
)

func GetAllCarsAPI() string {
	url := fmt.Sprintf("%s/search?api_key=%s", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY)

	return url
}

func GetCarByIdAPI(id string) string {
	url := fmt.Sprintf("%s/info?api_key=%s&auto_id=%s", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY, id)

	return url
}

func GetAllMarksAPI() string {
	url := fmt.Sprintf("%s/categories/:categoryId/marks?api_key=%s", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY)

	return url
}

func GetCarsByMarkAPI(mark int64) string {
	url := fmt.Sprintf("%s/search?api_key=%s&marka_id[0]=%d", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY, mark)

	return url
}
