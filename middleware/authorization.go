package middleware

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/model"
	"github.com/acm-uiuc/core/service"
)

func AuthorizeMark(svc *service.Service, marks []string) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*context.Context)
			if !ok {
				return ctx.String(http.StatusForbidden, "Invalid Context")
			}

			user, err := svc.User.GetUser(ctx.Username)
			if err != nil {
				return ctx.String(http.StatusForbidden, "Could Not Find User")
			}

			validMark := false
			for _, mark := range marks {
				if user.Mark == mark {
					validMark = true
				}
			}

			if !validMark {
				return ctx.String(http.StatusForbidden, "Invalid User Mark")
			}

			return next(ctx)
		}
	}
}

func AuthorizeCommittee(svc *service.Service, committees []string) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*context.Context)
			if !ok {
				return ctx.String(http.StatusForbidden, "Invalid Context")
			}

			validCommittee := false
			for _, committee := range committees {
				isMember, err := svc.Group.VerifyMembership(ctx.Username, model.GroupCommittees, committee)
				if err != nil {
					return ctx.String(http.StatusForbidden, "Failed Verifying Membership")
				}

				if isMember {
					validCommittee = true
				}
			}

			if !validCommittee {
				return ctx.String(http.StatusForbidden, "Invalid Group Membership")
			}

			return next(ctx)
		}
	}
}
