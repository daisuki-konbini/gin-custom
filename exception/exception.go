package exception

//Exception ...
type Exception struct {
	Code int
	Msg  string
}

func (e *Exception) Error() string {
	return e.Msg
}

//NewFrom ...
func NewFrom(originError error) *Exception {
	return &Exception{
		Code: InternalServerError.Code,
		Msg:  originError.Error(),
	}
}

var (
	OK = &Exception{Code: 0, Msg: "OK"}

	InternalServerError = &Exception{Code: 10001, Msg: "InternalServerError"}

	PayError = &Exception{Code: 20001, Msg: "pay error"}
)
