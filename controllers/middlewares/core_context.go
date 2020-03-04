package middlewares

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/controllers/context"
	"github.com/acm-uiuc/core/services"
)

func ContextExtender(svcs services.Services) (func(echo.HandlerFunc) echo.HandlerFunc) {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := ctx.Request().Header.Get("Authorization")

			username, err := svcs.Auth.Verify(token)
			if err != nil {
				return fmt.Errorf("failed to verify token: %w", err)
			}

			info, err := svcs.User.GetInfo(username)
			if err != nil {
				return fmt.Errorf("failed to get user info: %w", err)
			}

			memberships, err := svcs.Group.GetMemberships(username)
			if err != nil {
				return fmt.Errorf("failed to get memberships: %w", err)
			}

			coreContext := &context.CoreContext {
				Context: ctx,
				Username: username,
				Memberships: memberships,
				Mark: string(info.Mark),
			}

			return next(coreContext)
		}
	}
}
