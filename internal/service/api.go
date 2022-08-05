package server

import (
	"fmt"
	"ria-bot/configs"
)

func GetAllCarsAPI() string {
	return fmt.Sprintf("%s/search?api_key=%s", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY)
}

func GetCarByIdAPI(id string) string {
	return fmt.Sprintf("%s/info?api_key=%s&auto_id=%s", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY, id)
}

func GetAllMarksAPI() string {
	return fmt.Sprintf("%s/categories/:categoryId/marks?api_key=%s", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY)
}

func GetCarsByMarkAPI(mark int64) string {
	return fmt.Sprintf("%s/search?api_key=%s&marka_id[0]=%d", configs.RIA_DEV_BASE_URL, configs.RIA_API_KEY, mark)
}
