package controllers

import (
	"fmt"
	"net/http"

	"github.com/acm-uiuc/core/controllers/context"
	"github.com/acm-uiuc/core/services/user"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	GraduationYear int32 `json:"graduation_year"`
	Major string `json:"major"`
}

type CreateUserResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
}

func (controller *Controller) CreateUserController(ctx *context.CoreContext) error {
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

	err = controller.svcs.User.Create(user.UserData {
		Username: req.Username,
		FirstName: req.FirstName,
		LastName: req.LastName,
		GraduationYear: req.GraduationYear,
		Major: req.Major,
		Mark: string(user.MarkBasic),
	})
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
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	GraduationYear int32 `json:"graduation_year"`
	Major string `json:"major"`
	Mark string `json:"mark"`
}

func (controller *Controller) GetUserController(ctx *context.CoreContext) error {
	req := &GetUserRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &GetUserResponse {
				Success: false,
				Message: "Internal Error",
				Username: "",
				FirstName: "",
				LastName: "",
				GraduationYear: 0,
				Major: "",
				Mark: "",
			}),
		)
	}

	info, err := controller.svcs.User.GetInfo(req.Username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w",
			ctx.JSON(http.StatusBadRequest, &GetUserResponse {
				Success: false,
				Message: "Invalid Username",
				Username: "",
				FirstName: "",
				LastName: "",
				GraduationYear: 0,
				Major: "",
				Mark: "",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &GetUserResponse {
		Success: true,
		Message: "Successful User Retrieval",
		Username: info.Username,
		FirstName: info.FirstName,
		LastName: info.LastName,
		GraduationYear: info.GraduationYear,
		Major: info.Major,
		Mark: info.Mark,
	})
}

type GetUsersRequest struct {
	Username string `json:"username"`
}

type GetUsersResponse struct {
	Success bool `json:"success`
	Message string `json:"message"`
	Users []struct{
		Username string `json:"username"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		GraduationYear int32 `json:"graduation_year"`
		Major string `json:"major"`
		Mark string `json:"mark"`
	}
}

func (controller *Controller) GetUsersController(ctx *context.CoreContext) error {
	req := &GetUsersRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &GetUsersResponse {
				Success: false,
				Message: "Internal Error",
				Users: nil,
			}),
		)
	}

	infos, err := controller.svcs.User.GetInfos()
	if err != nil {
		return fmt.Errorf("failed to get user: %w",
			ctx.JSON(http.StatusBadRequest, &GetUsersResponse {
				Success: false,
				Message: "Invalid Username",
				Users: nil,
			}),
		)
	}

	resp := &GetUsersResponse {
		Success: true,
		Message: "Successful Users Retrieval",
		Users: nil,
	}

	for _, info := range infos {
		resp.Users = append(resp.Users, struct{
			Username string `json:"username"`
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
			GraduationYear int32 `json:"graduation_year"`
			Major string `json:"major"`
			Mark string `json:"mark"`
		} {
			Username: info.Username,
			FirstName: info.FirstName,
			LastName: info.LastName,
			GraduationYear: info.GraduationYear,
			Major: info.Major,
			Mark: info.Mark,
		})
	}

	return ctx.JSON(http.StatusOK, resp)
}
