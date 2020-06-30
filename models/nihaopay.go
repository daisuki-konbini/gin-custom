package models

import (
	"encoding/json"
	"gin/config"
	"gin/exception"
	"gin/helpers"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//HTTPClient ...
var HTTPClient = &http.Client{}

//Currency ...
type Currency string

const (
	USD Currency = "USD"
	JPY Currency = "JPY"
	HKD Currency = "HKD"
	GPB Currency = "GPB"
	EUR Currency = "EUR"
	CAD Currency = "CAD"
)

//vendor ...
type Vendor string

const (
	AliPay    Vendor = "alipay"
	WechatPay Vendor = "wechatpay"
	UnionPay  Vendor = "unionpay"
)

//AsyncMsg ...
type AsyncMsg struct {
	ID         string    `form:"id"`
	Amount     int       `form:"amount"`
	Currency   Currency  `form:"currency"`
	RMBAmount  int       `form:"rmb_amount"`
	Reference  string    `form:"reference"`
	SysReserve string    `form:"sys_reserve"`
	Status     string    `form:"status"`
	Time       time.Time `form:"time"`
	Note       string    `form:"note"`
	VerifySign string    `form:"verify_sign"`
}

//SecurePayPayload ...
type SecurePayPayload struct {
	Amount      int      `form:"amount"`
	Currency    Currency `form:"currency"`
	Vendor      Vendor   `form:"vendor"`
	Reference   string   `form:"reference"`
	IpnURL      string   `form:"ipn_url"`
	CallbackURL string   `form:"callback_url"`
}

//RefundPayload ...
type RefundPayload struct {
	Amount   int      `form:"amount"`
	Currency Currency `form:"currency"`
}

//RefundResponse ...
type RefundResponse struct {
	ID            string `form:"id"`
	Status        string `form:"status"`
	Refunded      bool   `form:"refunded"`
	TransactionID string `form:"transaction_id"`
}

//NihaopayError ...
type NihaopayError struct {
	code    int
	label   string
	message string
}

//SecurePay ...
func SecurePay(payload *SecurePayPayload) (*http.Response, error) {
	//TODO add log
	data := strings.NewReader(helpers.Struct2Map(payload).Encode())
	url := config.GetConfig().GetString("nihaopay.api") + "/transactions/securepay"
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	resp, err := HTTPClient.Do(headerWrapper(req))
	if err != nil {
		return nil, err
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		//TODO log error
		return nil, exception.PayError
	}
	return resp, nil
}

//Refund ...
func Refund(transactionID string, payload *RefundPayload) (res *RefundResponse, err error) {
	//TODO add log
	data := strings.NewReader(helpers.Struct2Map(payload).Encode())
	url := config.GetConfig().GetString("nihaopay.api") + "/transactions/" + transactionID + "/refund"
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	resp, err := HTTPClient.Do(headerWrapper(req))
	if err != nil || resp.StatusCode != http.StatusOK {
		//TODO log error
		return nil, exception.PayError
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, err
	}
	return
}

func headerWrapper(req *http.Request) *http.Request {
	req.Header.Set("Authorization", config.GetConfig().GetString("nihaopay.token"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}
