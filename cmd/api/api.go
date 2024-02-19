package api

import (
	"net/http"

	"github.com/anukuljoshi/goweb/config"
	"github.com/anukuljoshi/goweb/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Api struct {
	Server  *echo.Echo
	Cfg     *config.Config
	Service *services.Service
}

func NewApi(cfg *config.Config, service *services.Service) (*Api, error) {
	var api = &Api{}
	api.Cfg = cfg
	api.Service = service

	api.Server = echo.New()
	api.Server.HTTPErrorHandler = customErrorHandler
	// Root level middleware
	api.Server.Use(middleware.Logger())
	api.Server.Use(middleware.Recover())

	api.Server.Logger.SetLevel(log.DEBUG)

	api.Server.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, api.okResponse(map[string]string{
			"message": "pong",
		}))
	})

	usersGroup := api.Server.Group("/users")
	usersGroup.GET(
		"", api.GetUserList,
	)
	usersGroup.POST(
		"", api.PostUserCreate,
	)
	return api, nil
}

// Custom error handler
func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	// Format the response in JSON
	response := map[string]interface{}{
		"message": message,
		"data":    nil,
		"status":  code,
		"error":   true,
	}

	// Send the JSON response
	if !c.Response().Committed {
		c.JSON(code, response)
	}
}
