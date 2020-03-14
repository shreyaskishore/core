package docs

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
)

func Documentation(ctx *context.Context) error {
	// TODO: Add documentation for the endpoints
	return ctx.String(http.StatusOK, "Hello World!")
}
