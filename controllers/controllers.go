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

	controller.POST("/auth/login", ContextConverter(controller.LoginController))
	controller.POST("/auth/logout", ContextConverter(controller.LogoutController))
	controller.POST("/auth/verify", ContextConverter(controller.VerifyController))
	controller.POST("/auth/local", ContextConverter(controller.LocalAccountController))

	controller.POST("/user/create", ContextConverter(controller.CreateUserController))
	controller.POST("/user/mark", ContextConverter(controller.MarkUserController))
	controller.POST("/user/find", ContextConverter(controller.GetUserController))
	controller.POST("/user/all", ContextConverter(controller.GetUsersController))

	controller.POST("/group/memberships", ContextConverter(controller.MembershipsController))
	controller.POST("/group/groups", ContextConverter(controller.GroupsController))
	controller.POST("/group/verify", ContextConverter(controller.VerifyGroupController))

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
