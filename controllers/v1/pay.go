package v1

import (
	"fmt"
	"gin/config"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//PayController ...
type PayController struct{}

//SecurePay ...
func (p PayController) SecurePay(c *gin.Context) {
	payload := &models.SecurePayPayload{
		Amount:      1,
		Currency:    models.JPY,
		Vendor:      models.UnionPay,
		Reference:   uuid.New().String(),
		IpnURL:      config.GetConfig().GetString("nihaopay.ipn_url"),
		CallbackURL: config.GetConfig().GetString("nihaopay.callback_url"),
	}
	response, err := models.SecurePay(payload)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{
		"Content-Disposition": `inline`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

//AsyncMsg ...
func (p PayController) AsyncMsg(c *gin.Context) {
	var msg models.AsyncMsg
	c.ShouldBind(&msg)
	//TODO veify token
	fmt.Println(msg, msg.ID)
	c.String(http.StatusOK, "ok")
}

//Refund ...
func (p PayController) Refund(c *gin.Context) interface{} {
	payload := &models.RefundPayload{
		Amount:   1,
		Currency: models.JPY,
	}
	ret, err := models.Refund("202006300208120266191", payload)
	if err != nil {
		return err
	}
	return ret
}
