package context

import (
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	Username string
	LoggedIn bool
}
