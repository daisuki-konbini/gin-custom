package controllers

import (
	"gin/db"

	"github.com/gin-gonic/gin"
)

//HealthController ...
type HealthController struct{}

//Status ...
func (h HealthController) Status(c *gin.Context) interface{} {
	err := db.GetDB().DB().Ping()
	if err != nil {
		return err
	}

	return "ok1"
}
