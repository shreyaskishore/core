package site

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/service"
)

type SiteController struct {
	svc *service.Service
}

func New(svc *service.Service) *SiteController {
	return &SiteController{
		svc: svc,
	}
}

func (controller *SiteController) Home(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusOK, "home", params)
}
