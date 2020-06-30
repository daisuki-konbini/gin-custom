package responses

import (
	"gin/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandlerFunc ..
type HandlerFunc func(c *gin.Context) interface{}

//APIResponse ...
type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//APIWrapper ...
func APIWrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var res APIResponse
		ret := handler(c)
		switch v := ret.(type) {
		case *exception.Exception:
			res = APIResponse{
				Code: v.Code,
				Msg:  v.Msg,
				Data: nil,
			}
		case error:
			e := exception.NewFrom(v)
			res = APIResponse{
				Code: e.Code,
				Msg:  e.Msg,
				Data: nil,
			}
		default:
			res = APIResponse{
				Code: exception.OK.Code,
				Msg:  "",
				Data: v,
			}
		}
		c.JSON(http.StatusOK, res)
	}
}
