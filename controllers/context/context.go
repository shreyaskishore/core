package context

import (
	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/services/user"
)

type CoreContext struct {
	echo.Context
	Username string
	Memberships []string
	Mark string
}

func (ctx *CoreContext) HasMembership(group string) bool {
	for _, membership := range ctx.Memberships {
		if membership == group {
			return true
		}
	}

	return false;
}

func (ctx *CoreContext) HasValidToken() bool {
	return ctx.Username != ""
}

func (ctx *CoreContext) IsPaidMember() bool {
	return ctx.Mark == string(user.MarkPaid)
}

func (ctx *CoreContext) IsRecruiter() bool {
	return ctx.Mark == string(user.MarkRecruiter)
}
