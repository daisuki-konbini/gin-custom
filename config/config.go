package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var config *viper.Viper

//Init ...
func init() {
	gin.SetMode(gin.TestMode)
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(gin.Mode())
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

//GetConfig ...
func GetConfig() *viper.Viper {
	return config
}
