package login

import (
	"github.com/labstack/echo"

	"github.com/mariobac1/api_/domain/user"
)

// Route Login
func RouteUser(e *echo.Echo, usecase user.Storage) {
	h := newLogin(usecase)

	e.POST("/v1/login", h.login)
	e.POST("/v1/sign-up", h.create)
}
