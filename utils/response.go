package utils

type Response struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Error  interface{} `json:"error,omitempty"` //for errors that occur even if request is successful
	Data   interface{} `json:"data,omitempty"`
}

func SuccessResponse(code int, data interface{}) Response {
	res := ResponseMessage(code, "success", nil, data)
	return res
}

func ErrorResponse(code int, status string, err interface{}, data interface{}) Response {
	res := ResponseMessage(code, status, err, data)
	return res
}

func ResponseMessage(code int, status string, err interface{}, data interface{}) Response {
	res := Response{
		Code:   code,
		Status: status,
		Error:  err,
		Data:   data,
	}
	return res
}
