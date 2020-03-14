package docs

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
)

func Documentation(ctx *context.Context) error {
	return ctx.String(http.StatusOK, "Hello World!")
}
