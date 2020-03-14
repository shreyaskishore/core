package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/model"
	"github.com/acm-uiuc/core/service"
)

type AuthorizeMatchParameters struct {
	Marks      []string
	Committees []string
}

func AuthorizeMatchAny(svc *service.Service, match AuthorizeMatchParameters) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*context.Context)
			if !ok {
				return ctx.String(http.StatusForbidden, "Invalid Context")
			}

			isMarkMatch, err := hasMarkMatch(svc, ctx.Username, match.Marks)
			if err != nil {
				return ctx.String(http.StatusForbidden, "Failed Mark Match")
			}

			isCommitteeMatch, err := hasCommitteeMatch(svc, ctx.Username, match.Committees)
			if err != nil {
				return ctx.String(http.StatusForbidden, "Failed Committee Match")
			}

			if !isMarkMatch && !isCommitteeMatch {
				return ctx.String(http.StatusForbidden, "Invalid Authorization Matches")
			}

			return next(ctx)
		}
	}
}

func hasMarkMatch(svc *service.Service, username string, marks []string) (bool, error) {
	user, err := svc.User.GetUser(username)
	if err != nil {
		return false, fmt.Errorf("failed to find user: %w", err)
	}

	for _, mark := range marks {
		if user.Mark == mark {
			return true, nil
		}
	}

	return false, nil
}

func hasCommitteeMatch(svc *service.Service, username string, committees []string) (bool, error) {
	for _, committee := range committees {
		isMember, err := svc.Group.VerifyMembership(username, model.GroupCommittees, committee)
		if err != nil {
			return false, fmt.Errorf("failed verifying membership: %w", err)
		}

		if isMember {
			return true, nil
		}
	}

	return false, nil
}
