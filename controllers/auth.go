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

type LogoutRequest struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func (controller *Controller) LogoutController(ctx *context.CoreContext) error {
	req := &LogoutRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &LogoutResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	err = controller.svcs.Auth.Logout(req.Token)
	if err != nil {
		return fmt.Errorf("failed to logout: %w",
			ctx.JSON(http.StatusBadRequest, &LogoutResponse {
				Success: false,
				Message: "Invalid Token",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &LogoutResponse {
		Success: true,
		Message: "Successful Logout",
	})
}

type VerifyRequest struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	Success bool `json:"success"`
	Username string `json:"username"`
	Message string `json:"message"`
}

func (controller *Controller) VerifyController(ctx *context.CoreContext) error {
	req := &VerifyRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &VerifyResponse {
				Success: false,
				Username: "",
				Message: "Internal Error",
			}),
		)
	}

	username, err := controller.svcs.Auth.Verify(req.Token)
	if err != nil {
		return fmt.Errorf("failed to verify: %w",
			ctx.JSON(http.StatusBadRequest, &VerifyResponse {
				Success: false,
				Username: "",
				Message: "Invalid Token",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &VerifyResponse {
		Success: true,
		Username: username,
		Message: "Successful Verification",
	})
}

type LocalAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"username"`
}

type LocalAccountResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func (controller *Controller) LocalAccountController(ctx *context.CoreContext) error {
	req := &LocalAccountRequest {}
	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("failed to bind: %w",
			ctx.JSON(http.StatusInternalServerError, &LocalAccountResponse {
				Success: false,
				Message: "Internal Error",
			}),
		)
	}

	err = controller.svcs.Auth.CreateLocalAccount(req.Username, req.Password)
	if err != nil {
		return fmt.Errorf("failed to create local account: %w",
			ctx.JSON(http.StatusBadRequest, &LocalAccountResponse {
				Success: false,
				Message: "Invalid Token",
			}),
		)
	}

	return ctx.JSON(http.StatusOK, &LocalAccountResponse {
		Success: true,
		Message: "Successful Local Account Creation",
	})
}
