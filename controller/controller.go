package controller

import (
	"fmt"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/middleware"
	"github.com/acm-uiuc/core/service"

	_ "github.com/acm-uiuc/core/controller/auth"
	"github.com/acm-uiuc/core/controller/docs"
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

	docsController := docs.New(controller.svc)

	controller.Use(echoMiddleware.Logger())
	controller.Use(echoMiddleware.Recover())
	controller.Use(middleware.Context(controller.svc))

	controller.GET("/api", ContextConverter(docsController.Documentation))
	// TODO: Register routes

	return controller, nil
}

func ContextConverter(handler func(*context.Context) error) func(echo.Context) error {
	return func(c echo.Context) error {
		ctx, ok := c.(*context.Context)
		if !ok {
			return fmt.Errorf("failed to convert context")
		}
		return handler(ctx)
	}
}
