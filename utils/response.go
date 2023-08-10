package utils

type Response struct {
	Code       int         `json:"code,omitempty"`
	Status     string      `json:"status,omitempty"`
	Error      interface{} `json:"error,omitempty"` //for errors that occur even if request is successful
	ErrMessage string      `json:"errMessage,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func SuccessResponse(code int, data interface{}) Response {
	res := SuccessResponseMessage(code, "success", nil, data)
	return res
}

func ErrorResponse(code int, status string, err interface{}, errMessage string) Response {
	res := ErrorResponseMessage(code, status, err, errMessage, nil)
	return res
}

func SuccessResponseMessage(code int, status string, err interface{}, data interface{}) Response {
	res := Response{
		Code:   code,
		Status: status,
		Error:  err,
		Data:   data,
	}
	return res
}

func ErrorResponseMessage(code int, status string, err interface{}, errMessage string, data interface{}) Response {
	res := Response{
		Code:       code,
		Status:     status,
		Error:      err,
		ErrMessage: errMessage,
		Data:       data,
	}
	return res
}
