package docs

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/service"
)

type DocsController struct {
	svc *service.Service
}

func New(svc *service.Service) *DocsController {
	return &DocsController{
		svc: svc,
	}
}

func (controller *DocsController) Documentation(ctx *context.Context) error {
	// TODO: Add documentation for the endpoints
	return ctx.String(http.StatusOK, "Hello World!")
}
