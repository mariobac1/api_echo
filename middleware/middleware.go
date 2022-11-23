package middleware

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/mariobac1/api_/authorization"
)

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			//Non authorized
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Non authorized"})
		}
		return f(c)
	}
}
