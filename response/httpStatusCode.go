package response

const (
	ErrCodeSuccess      = 2001
	ErrCodeParamInvalid = 2003
	ErrCodeInvalidToken = 3001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Param Invalid",
	ErrCodeInvalidToken: "Invalid Token",
}
