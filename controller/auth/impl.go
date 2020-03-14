package auth

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/service"
)

type AuthController struct {
	svc *service.Service
}

func New(svc *service.Service) *AuthController {
	return &AuthController{
		svc: svc,
	}
}

func (controller *AuthController) GetOAuthRedirect(ctx *context.Context) error {
	provider := ctx.Param("provider")

	uri, err := controller.svc.Auth.GetOAuthRedirect(provider)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid Provider")
	}

	return ctx.Redirect(http.StatusFound, uri)
}

func (controller *AuthController) GetOAuthRedirectLanding(ctx *context.Context) error {
	return ctx.Render(http.StatusOK, "redirect", nil)
}

func (controller *AuthController) GetToken(ctx *context.Context) error {
	provider := ctx.Param("provider")

	req := struct {
		Code string `json:"code"`
	}{}

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Bind")
	}

	token, err := controller.svc.Auth.Authorize(provider, req.Code)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Token Generation")
	}

	return ctx.JSON(http.StatusOK, token)
}
