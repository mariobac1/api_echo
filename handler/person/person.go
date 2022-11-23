package person

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/mariobac1/api_/domain/person"
	"github.com/mariobac1/api_/models"
)

type handler struct {
	usecase person.Storage
}

func newHandler(usecase person.Storage) handler {
	return handler{usecase}
}

// handler Create a person
func (h *handler) create(c echo.Context) error {
	data := models.Person{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, "The structure is wrong", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err = h.usecase.Create(&data)
	if err != nil {
		resp := NewResponse(Error, "An issue occurs when try create a person", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := NewResponse(Message, "Person created Ok", nil)
	return c.JSON(http.StatusOK, resp)
}

// handler GetAll persons
func (h *handler) getAll(c echo.Context) error {
	data, err := h.usecase.GetAll()
	if err != nil {
		resp := NewResponse(Error, "An issue occurs when try get all persons", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, "OK", data)
	return c.JSON(http.StatusOK, resp)

}

// Handler Update
func (h *handler) update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID < 1 {
		resp := NewResponse(Error, "The ID will be positive", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	data := models.Person{}
	err = c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, "The structure is wrong", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	data.ID = uint(ID)
	data.UpdatedAt = time.Now()

	err = h.usecase.Update(&data)
	if errors.Is(err, models.ErrIDPersonDoesNotExists) {
		resp := NewResponse(Error, "This ID doesn't exist", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if err != nil {
		resp := NewResponse(Error, "An issue occurs when try update a person", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := NewResponse(Message, "Person update ok", nil)
	return c.JSON(http.StatusOK, resp)
}

// handler byid
func (h *handler) getById(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID < 1 {
		resp := NewResponse(Error, "The ID will be positive", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	id := uint(ID)
	data := &models.Person{}
	data, err = h.usecase.GetByID(id)
	fmt.Printf("el error es: %v", err)
	if errors.Is(err, models.ErrIDPersonDoesNotExists) {
		resp := NewResponse(Error, "This ID doesn't exist", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if err != nil {
		resp := NewResponse(Error, "An issue occurs when try get a person", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := NewResponse(Message, "Person update ok", data)
	return c.JSON(http.StatusOK, resp)
}

// Handler Delete person
func (h *handler) delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID < 1 {
		resp := NewResponse(Error, "The ID will be positive", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = h.usecase.Delete(uint(ID))
	if errors.Is(err, models.ErrIDPersonDoesNotExists) {
		resp := NewResponse(Error, "This ID doesn't exist", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if err != nil {
		resp := NewResponse(Error, "An issue occurs when try update a person", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, "Person delete ok", nil)
	return c.JSON(http.StatusOK, resp)
}
