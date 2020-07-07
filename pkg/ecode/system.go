package ecode

var (
	OK = New(0, "OK")
	//System[10000,20000) ...
	InternalServerError = New(10001, "InternalServerError")
	TokenExpired        = New(10002, "token expired")
	TokenErr            = New(10003, "can not handle the token")

	//pay[20000,30000)
	PayError = New(20001, "pay error")
)
