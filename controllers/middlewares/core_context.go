package middlewares

import (
	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/controllers/context"
)

func ContextExtender(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		coreContext := &context.CoreContext{ctx}
		return next(coreContext)
	}
}
