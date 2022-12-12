package common

import "net/http"

type successRes struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Paging     any    `json:"paging,omitempty"`
	Filter     any    `json:"filter,omitempty"`
	Success    bool   `json:"success"`
}

func NewSuccessResponse(statusCode int, message string, data, paging, filter any) *successRes {
	return &successRes{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Paging:     paging,
		Filter:     filter,
		Success:    true,
	}
}
func SimpleSuccessResponse(data any) *successRes {
	return NewSuccessResponse(
		http.StatusOK, "success",
		data,
		nil,
		nil,
	)
}

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
	Success    bool   `json:"success"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
		Success:    false,
	}
}
