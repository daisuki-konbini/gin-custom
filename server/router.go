package server

import (
	"gin/controllers"
	v1 "gin/controllers/v1"
	"gin/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandlerFunc ..
type HandlerFunc func(c *gin.Context) interface{}

//Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var res Response
		ret := handler(c)
		switch v := ret.(type) {
		case *exception.Exception:
			res = Response{
				Code: v.Code,
				Msg:  v.Msg,
				Data: nil,
			}
		case error:
			e := exception.NewFrom(v)
			res = Response{
				Code: e.Code,
				Msg:  e.Msg,
				Data: nil,
			}
		default:
			res = Response{
				Code: exception.OK.Code,
				Msg:  "",
				Data: v,
			}
		}
		c.JSON(http.StatusOK, res)
	}
}

//SetRouter ...
func SetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//health check
	r.GET("/health", wrapper((controllers.HealthController{}).Status))

	r.GET("/user", (controllers.UserController{}).Test)

	g1 := r.Group("/v1")
	{
		g1.POST("/register", (v1.RegisterController{}).RegisterWithEmail)
	}

	return r
}
