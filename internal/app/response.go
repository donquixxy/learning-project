package app

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type FailedResponse struct {
	Message string `json:"message"`
}

func ResponseFailed(errMsg string) *FailedResponse {
	return &FailedResponse{
		Message: errMsg,
	}
}

func ResponseSuccess(data any, msg string) *SuccessResponse {
	return &SuccessResponse{
		Message: msg,
		Data:    data,
	}
}
