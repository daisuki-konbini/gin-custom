package ecode

//Ecode ...
type Ecode struct {
	Code int
	Msg  string
}

func (e *Ecode) Error() string {
	return e.Msg
}

//New ...
func New(code int, msg string) *Ecode {
	return &Ecode{
		Code: code,
		Msg:  msg,
	}
}
