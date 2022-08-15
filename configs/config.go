package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	RiaApiKey      string `mapstructure:"RIA_API_KEY"`
	RiaBaseUrl     string `mapstructure:"RIA_BASE_URL"`
	RiaDevBaseUrl  string `mapstructure:"RIA_DEV_BASE_URL"`
	TelegramApiKey string `mapstructure:"TELEGRAM_API_KEY"`
}

func InitViper(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
