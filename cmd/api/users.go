package api

import (
	"github.com/anukuljoshi/goweb/models"
	"github.com/labstack/echo/v4"
)

func (a *Api) GetUserList(c echo.Context) error {
	user_list, err := a.Service.User.ListUsers()
	if err != nil {
		response := a.serverErrorResponse(err)
		return c.JSON(response.Status, response)
	}
	response := a.okResponse(user_list)
	return c.JSON(response.Status, response)
}

func (a *Api) PostUserCreate(c echo.Context) error {
	user_body := new(models.CreateUserParams)
	err := c.Bind(user_body)
	if err != nil {
		response := a.badRequestResponse(err)
		return c.JSON(response.Status, response)
	}
	user, err := a.Service.User.CreateUser(
		models.CreateUserParams(*user_body),
	)
	if err != nil {
		response := a.serverErrorResponse(err)
		return c.JSON(response.Status, response)
	}
	response := a.createdResponse(user)
	return c.JSON(response.Status, response)
}
