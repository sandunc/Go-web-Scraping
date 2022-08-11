package utils

import (
	"Go-web-Scraping/consts"

	"log"

	"github.com/spf13/viper"
)

func init() {
	log.Printf("Starting APP................")
	viper.SetConfigFile(consts.ENV_FILE)
	viper.AddConfigPath(consts.ENV_FILE_DIRECTORY)
	viper.AddConfigPath(consts.TEST_ENV_FILE_DIRECTORY)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error occurred while reading env file, might fallback to OS env config")
	}
	viper.AutomaticEnv()
}

func GetEnvVar(name string) string {
	if !viper.IsSet(name) {
		log.Fatal("nvironment variable is not set")
		return ""
	}
	value := viper.GetString(name)
	return value
}
