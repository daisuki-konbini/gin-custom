package orm

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Config ...
type Config struct {
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration //
}

//New ...
func New(c *Config) *gorm.DB {
	db, err := gorm.Open("postgres", c.DSN)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetConnMaxLifetime(c.IdleTimeout)
	return db
}
