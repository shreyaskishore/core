package controllers

import (
	"net/http"

	"github.com/acm-uiuc/core/controllers/context"
)

func LoginController(ctx *context.CoreContext) error {
	return ctx.String(http.StatusOK, "Hello World!")
}
