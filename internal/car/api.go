package server

import (
	"fmt"

	configs "ria-bot/configs"
)

func GetAllCarsAPI(cfg *configs.Config) string {
	return fmt.Sprintf("%s/search?api_key=%s", cfg.RiaDevBaseUrl, cfg.RiaApiKey)
}

func GetCarByIdAPI(id string, cfg *configs.Config) string {
	return fmt.Sprintf("%s/info?api_key=%s&auto_id=%s", cfg.RiaDevBaseUrl, cfg.RiaApiKey, id)
}

func GetAllMarksAPI(cfg *configs.Config) string {
	return fmt.Sprintf("%s/categories/:categoryId/marks?api_key=%s", cfg.RiaDevBaseUrl, cfg.RiaApiKey)
}

func GetCarsByMarkAPI(mark int64, cfg *configs.Config) string {
	return fmt.Sprintf("%s/search?api_key=%s&marka_id[0]=%d", cfg.RiaDevBaseUrl, cfg.RiaApiKey, mark)
}
