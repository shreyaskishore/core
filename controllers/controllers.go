package controllers

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/acm-uiuc/core/services"
	"github.com/acm-uiuc/core/controllers/middlewares"
	"github.com/acm-uiuc/core/controllers/context"
)

type Controller struct {
	*echo.Echo
	svcs services.Services
}

func New(svcs services.Services) Controller {
	controller := Controller {
		echo.New(),
		svcs,
	}

	controller.Use(middleware.Logger())
	controller.Use(middleware.Recover())
	controller.Use(middlewares.ContextExtender)
	// TODO: Add more middlewares

	controller.POST("/auth/login", ContextConverter(controller.LoginController))
	// TODO: Add more routes

	return controller
}

func ContextConverter(controller func(*context.CoreContext) error) func(echo.Context) error {
	return func(ctx echo.Context) error {
		coreContext, ok := ctx.(*context.CoreContext)
		if !ok {
			return fmt.Errorf("failed to convert context")
		}
		return controller(coreContext)
	}
}
