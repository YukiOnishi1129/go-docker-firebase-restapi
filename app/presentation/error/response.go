package error

import "github.com/harakeishi/gats"

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrorResponse struct {
	Status int `json:"status"`
	Error  Error
}

func NewErrorResponse(status int, errMessage interface{}) ErrorResponse {
	message, _ := gats.ToString(errMessage)
	return ErrorResponse{
		Status: status,
		Error: Error{
			Message: message,
			Code:    status,
		},
	}
}
