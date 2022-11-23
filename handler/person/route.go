package person

import (
	"github.com/labstack/echo"

	"github.com/mariobac1/api_/domain/person"
	"github.com/mariobac1/api_/middleware"
)

func RoutePerson(e echo.Echo, usecase person.Storage) {
	h := newHandler(usecase)
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.GET("", h.getAll)
	person.GET("/:id", h.getById)
	person.PUT("/:id", h.update)
	person.DELETE("/:id", h.delete)
}
