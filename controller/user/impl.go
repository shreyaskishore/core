package user

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/model"
	"github.com/acm-uiuc/core/service"
)

type UserController struct {
	svc *service.Service
}

func New(svc *service.Service) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (controller *UserController) GetUser(ctx *context.Context) error {
	user, err := controller.svc.User.GetUser(ctx.Username)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed User Lookup")
	}

	return ctx.JSON(http.StatusOK, user)
}

func (controller *UserController) CreateUser(ctx *context.Context) error {
	req := model.User{}

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Bind")
	}

	req.Mark = model.UserMarkBasic

	err = controller.svc.User.CreateUser(req)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed User Creation")
	}

	return ctx.JSON(http.StatusOK, struct{}{})
}

func (controller *UserController) GetUsers(ctx *context.Context) error {
	users, err := controller.svc.User.GetUsers()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Users Lookup")
	}

	return ctx.JSON(http.StatusOK, users)
}

func (controller *UserController) MarkUser(ctx *context.Context) error {
	req := struct {
		Username string `json:"username"`
		Mark     string `json:"mark"`
	}{}

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Bind")
	}

	err = controller.svc.User.MarkUser(req.Username, req.Mark)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed User Marking")
	}

	return ctx.JSON(http.StatusOK, struct{}{})
}

func (controller *UserController) DeleteUser(ctx *context.Context) error {
	req := model.User{}

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Bind")
	}

	err = controller.svc.User.DeleteUser(req.Username)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed User Deletion")
	}

	return ctx.JSON(http.StatusOK, struct{}{})
}
