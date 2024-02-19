package api

import (
	"net/http"
)

// create base success response structure
func (a *Api) successResponse(message string, data any, status int) Response {
	return Response{
		Message: message,
		Data:    data,
		Status:  status,
		Error:   false,
	}
}

func (a *Api) okResponse(data any) Response {
	return a.successResponse(
		http.StatusText(http.StatusOK),
		data,
		http.StatusOK,
	)
}

func (a *Api) createdResponse(data any) Response {
	return a.successResponse(
		http.StatusText(http.StatusCreated),
		data,
		http.StatusCreated,
	)
}
