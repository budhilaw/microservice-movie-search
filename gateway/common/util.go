package common

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
}

func APIResponse(message string, data interface{}) Response {
	meta := Meta{
		Message: message,
	}

	res := Response{
		Meta: meta,
		Data: data,
	}

	return res
}
