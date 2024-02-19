package api

import (
	"net/http"
)

// create base error response structure
func (a *Api) errorResponse(
	message string,
	status int,
) Response {
	return Response{
		Message: message,
		Data:    nil,
		Status:  status,
		Error:   true,
	}
}

func (a *Api) serverErrorResponse(err error) Response {
	a.Server.Logger.Error(err.Error())
	return a.errorResponse(
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

func (a *Api) notFoundResponse() Response {
	return a.errorResponse(
		http.StatusText(http.StatusNotFound),
		http.StatusNotFound,
	)
}

func (a *Api) badRequestResponse(err error) Response {
	return a.errorResponse(
		err.Error(),
		http.StatusBadRequest,
	)
}
