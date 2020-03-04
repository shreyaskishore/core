package controllers

import (
	"fmt"
	"net/http"

	"github.com/acm-uiuc/core/controllers/context"
)

type LoginRequest struct {
	Username string `json:"username`
	Password string `json:"password`
}

type LoginResponse struct {
	Success bool `json:"success`
	Token string `json:"token"`
	Message string `json:"message"`
}

func (controller *Controller) LoginController(ctx *context.CoreContext) error {
	req := &LoginRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &LoginResponse {
				Success: false,
				Token: "",
				Message: "Internal Error",
			}),
		)
	}

	token, err := controller.svcs.Auth.Login(req.Username, req.Password)
	if err != nil {
		return fmt.Errorf("failed to login: %w",
			ctx.JSON(http.StatusUnauthorized, &LoginResponse {
				Success: false,
				Token: "",
				Message: "Invalid Login Details",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &LoginResponse {
		Success: true,
		Token: token,
		Message: "Successful Login",
	})
}
