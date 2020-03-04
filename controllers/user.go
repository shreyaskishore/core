package controllers

import (
	"fmt"
	"net/http"

	"github.com/acm-uiuc/core/controllers/context"
	"github.com/acm-uiuc/core/services/user"
	"github.com/acm-uiuc/core/services/group"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	user.UserData
}

type CreateUserResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
}

func (controller *Controller) CreateUserController(ctx *context.CoreContext) error {
	if !ctx.HasValidToken() {
		return fmt.Errorf("unauthorized: %w",
			ctx.JSON(http.StatusUnauthorized, &CreateUserResponse {
				Success: false,
				Message: "Invalid Authorization",
			}),
		)
	}

	req := &CreateUserRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &CreateUserResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	req.Mark = string(user.MarkBasic)

	err = controller.svcs.User.Create(req.UserData)
	if err != nil {
		return fmt.Errorf("failed to create user: %w",
			ctx.JSON(http.StatusBadRequest, &CreateUserResponse {
				Success: false,
				Message: "Invalid User Details",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &CreateUserResponse {
		Success: true,
		Message: "Successful User Creation",
	})
}

type MarkUserRequest struct {
	Username string `json:"username"`
	Mark string `json:"mark"`
}

type MarkUserResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
}

func (controller *Controller) MarkUserController(ctx *context.CoreContext) error {
	if !ctx.HasMembership(group.GroupTop4) {
		return fmt.Errorf("unauthorized: %w",
			ctx.JSON(http.StatusUnauthorized, &MarkUserResponse {
				Success: false,
				Message: "Invalid Authorization",
			}),
		)
	}

	req := &MarkUserRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &MarkUserResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	err = controller.svcs.User.Mark(req.Username, user.Mark(req.Mark))
	if err != nil {
		return fmt.Errorf("failed to mark user: %w",
			ctx.JSON(http.StatusBadRequest, &MarkUserResponse {
				Success: false,
				Message: "Invalid User Mark",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &MarkUserResponse {
		Success: true,
		Message: "Successful User Marking",
	})
}

type GetUserRequest struct {
	Username string `json:"username"`
}

type GetUserResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
	user.UserData
}

func (controller *Controller) GetUserController(ctx *context.CoreContext) error {
	req := &GetUserRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &GetUserResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	if !(ctx.IsRecruiter() || (ctx.HasValidToken() && ctx.Username == req.Username)) {
		return fmt.Errorf("unauthorized: %w",
			ctx.JSON(http.StatusUnauthorized, &MarkUserResponse {
				Success: false,
				Message: "Invalid Authorization",
			}),
		)
	}

	info, err := controller.svcs.User.GetInfo(req.Username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w",
			ctx.JSON(http.StatusBadRequest, &GetUserResponse {
				Success: false,
				Message: "Invalid Username",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &GetUserResponse {
		Success: true,
		Message: "Successful User Retrieval",
		UserData: info,
	})
}

type GetUsersRequest struct {
	Username string `json:"username"`
}

type GetUsersResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
	Users []user.UserData
}

func (controller *Controller) GetUsersController(ctx *context.CoreContext) error {
	if !(ctx.IsRecruiter() || ctx.HasMembership(group.GroupTop4) || ctx.HasMembership(group.GroupCorporate)) {
		return fmt.Errorf("unauthorized: %w",
			ctx.JSON(http.StatusUnauthorized, &MarkUserResponse {
				Success: false,
				Message: "Invalid Authorization",
			}),
		)
	}

	req := &GetUsersRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &GetUsersResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	infos, err := controller.svcs.User.GetInfos()
	if err != nil {
		return fmt.Errorf("failed to get user: %w",
			ctx.JSON(http.StatusBadRequest, &GetUsersResponse {
				Success: false,
				Message: "Invalid Username",
			}),
		)
	}


	return ctx.JSON(http.StatusOK, &GetUsersResponse {
		Success: true,
		Message: "Successful Users Retrieval",
		Users: infos,
	})
}
