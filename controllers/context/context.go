package context

import (
	"github.com/labstack/echo"
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
