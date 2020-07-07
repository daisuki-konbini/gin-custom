package http

import (
	"gin-custom/pkg/ecode"
	"net/http"

	"github.com/gin-gonic/gin"
)

//handlerFunc ..
type handlerFunc func(c *gin.Context) interface{}

//response ...
type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//wrapper ...
func wrapper(handler handlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var res response
		ret := handler(c)
		switch v := ret.(type) {
		case *ecode.Ecode:
			res = response{
				Code: v.Code,
				Msg:  v.Msg,
				Data: nil,
			}
		case error:
			e := ecode.New(ecode.InternalServerError.Code, v.Error())
			res = response{
				Code: e.Code,
				Msg:  e.Msg,
				Data: nil,
			}
		default:
			res = response{
				Code: ecode.OK.Code,
				Msg:  "",
				Data: v,
			}
		}
		c.JSON(http.StatusOK, res)
	}
}
