package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//New ...
func New() *viper.Viper {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(gin.Mode())
	config.AddConfigPath("../../configs/")
	config.AddConfigPath("configs/")
	viper.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return config
}
