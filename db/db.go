package db

import (
	"fmt"
	"gin/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//Init ...
func init() {
	var err error
	conf := config.GetConfig().Sub("db")
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			conf.GetString("host"),
			conf.GetString("port"),
			conf.GetString("user"),
			conf.GetString("dbname"),
			conf.GetString("password"),
		))
	if err != nil {
		//TODO panic
		log.Fatal(err)
	}
	db.DB().SetMaxOpenConns(conf.GetInt("maxcons"))
	db.DB().SetMaxIdleConns(conf.GetInt("maxidlecons"))
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
