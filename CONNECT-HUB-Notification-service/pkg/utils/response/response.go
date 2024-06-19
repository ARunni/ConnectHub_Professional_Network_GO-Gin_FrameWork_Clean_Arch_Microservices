package response

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClientResponse(statuscode int, message string, data interface{}, err interface{}) Response {
	return Response{
		StatusCode: statuscode,
		Message:    message,
		Data:       data,
		Error:      err,
	}
}
