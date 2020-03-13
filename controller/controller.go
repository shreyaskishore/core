package controller

import (
	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/service"

	_ "github.com/acm-uiuc/core/controller/auth"
	_ "github.com/acm-uiuc/core/controller/group"
	_ "github.com/acm-uiuc/core/controller/user"
)

type Controller struct {
	*echo.Echo
	svc service.Service
}

func New(svc service.Service) (*Controller, error) {
	controller := &Controller{
		Echo: echo.New(),
		svc:  svc,
	}

	// TODO: Inject middlewares

	// TODO: Register routes

	return controller, nil
}
