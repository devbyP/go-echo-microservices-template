package handlers

type errorResponse struct {
	Code   int               `json:"code"`
	Reason string            `json:"reason,omitempty"`
	Fields map[string]string `json:"fields,omitempty"`
}

type responseBody struct {
	Result  string         `json:"result"`
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *errorResponse `json:"error,omitempty"`
}

func newResponseBody(mess string, d interface{}) *responseBody {
	return &responseBody{
		Result:  "success",
		Message: mess,
		Data:    d,
	}
}

func newErrorResponse(code int, mess string, err error) *responseBody {
	error := &errorResponse{
		Code:   code,
		Reason: err.Error(),
	}
	res := &responseBody{
		Result:  "fail",
		Message: mess,
	}
	res.Error = error
	return res
}
