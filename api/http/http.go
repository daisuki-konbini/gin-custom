package http

import (
	"gin-custom/internal/app/account"
	"gin-custom/internal/app/health"
	"gin-custom/pkg/config"
	"gin-custom/pkg/database/orm"
	"gin-custom/pkg/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	conf *viper.Viper
	db   *gorm.DB

	//services
	healthSrv  health.Service
	accountSrv account.Service
)

//Init ...
func Init() {
	//config
	conf = config.New()

	//orm
	db = orm.New(&orm.Config{
		DSN:         conf.GetString("db.dsn"),
		Active:      conf.GetInt("db.active"),
		Idle:        conf.GetInt("db.idle"),
		IdleTimeout: time.Duration(conf.GetInt64("db.idle_timeout")),
	})

	//services
	healthSrv = health.NewService(db)
	accountSrv = account.NewService(db, conf)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cros())

	setRouter(r)

	r.Run(":8080")

}

func setRouter(r *gin.Engine) {
	//health check
	r.GET("/health", healthSrv.Check)

	v := r.Group("/v1")
	{
		v.POST("/register", wrapper(accountSrv.Register))
		v.POST("/login", wrapper(accountSrv.Login))
	}
}
