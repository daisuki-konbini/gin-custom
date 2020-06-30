package server

import (
	c "gin/controllers"
	v1 "gin/controllers/v1"
	res "gin/responses"

	"github.com/gin-gonic/gin"
)

//SetRouter ...
func SetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//health check
	r.GET("/health", res.APIWrapper((c.HealthController{}).Status))

	r.GET("/user", (c.UserController{}).Test)

	g1 := r.Group("/v1")
	{
		g1.POST("/register", (v1.RegisterController{}).RegisterWithEmail)
		g1.GET("/pay", (v1.PayController{}).SecurePay)
		g1.POST("/pay/callback", (v1.PayController{}).AsyncMsg)
		g1.POST("/pay/refund", res.APIWrapper((v1.PayController{}).Refund))
	}

	return r
}
