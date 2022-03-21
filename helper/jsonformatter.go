package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(code int, data interface{}, message, status string) Response {
	metaResponse := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta: metaResponse,
		Data: data,
	}
	return jsonResponse
}
