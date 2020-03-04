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
