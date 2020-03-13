package controller

import (
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"

	"github.com/acm-uiuc/core/middleware"
	"github.com/acm-uiuc/core/service"

	_ "github.com/acm-uiuc/core/controller/auth"
	_ "github.com/acm-uiuc/core/controller/group"
	_ "github.com/acm-uiuc/core/controller/user"
)

type Controller struct {
	*echo.Echo
	svc *service.Service
}

func New(svc *service.Service) (*Controller, error) {
	controller := &Controller{
		Echo: echo.New(),
		svc:  svc,
	}

	controller.Use(echoMiddleware.Logger())
	controller.Use(echoMiddleware.Recover())
	controller.Use(middleware.Context(controller.svc))

	// TODO: Inject middlewares

	// TODO: Register routes

	return controller, nil
}
