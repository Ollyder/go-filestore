package data

type Response struct {
	Status bool        `json:"status"`
	Reason string      `json:"reason"`
	Data   interface{} `json:"data"`
}

func Success(data interface{}) Response {
	resp := Response{
		Status: true,
		Data:   data,
	}
	return resp
}

func Fail(reason string) Response {
	resp := Response{
		Status: false,
		Reason: reason,
	}
	return resp
}
