package middlewares

import (
	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/controllers/context"
	"github.com/acm-uiuc/core/services"
	"github.com/acm-uiuc/core/services/user"
)

func ContextExtender(svcs services.Services) (func(echo.HandlerFunc) echo.HandlerFunc) {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := ctx.Request().Header.Get("Authorization")

			username := ""
			foundUsername, err := svcs.Auth.Verify(token)
			if err == nil {
				username = foundUsername
			}

			mark := string(user.MarkBasic)
			if username != "" {
				info, err := svcs.User.GetInfo(username)
				if err == nil {
					mark = info.Mark
				}
			}

			memberships := []string{}
			if username != "" {
				foundMemberships, err := svcs.Group.GetMemberships(username)
				if err == nil {
					memberships = foundMemberships
				}
			}

			coreContext := &context.CoreContext {
				Context: ctx,
				Username: username,
				Memberships: memberships,
				Mark: mark,
			}

			return next(coreContext)
		}
	}
}
