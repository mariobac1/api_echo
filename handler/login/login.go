package login

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/mariobac1/api_/authorization"
	"github.com/mariobac1/api_/domain/user"
	per "github.com/mariobac1/api_/handler/person"
	"github.com/mariobac1/api_/models"
)

const (
	Error   = "error"
	Message = "message"
)

type login struct {
	storage user.Storage
}

func newLogin(s user.Storage) login {
	return login{s}
}

// handler Create a person
func (l *login) login(c echo.Context) error {
	data := models.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := per.NewResponse(Error, "struct no valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	_, valid, _ := l.storage.GetByEmail(data.Email, data.Password)

	if !valid {
		resp := per.NewResponse(Error, "user or password not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := per.NewResponse(Error, "We can't make a new token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	resp := per.NewResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

// handler signUp a User
func (l *login) create(c echo.Context) error {
	data := models.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := per.NewResponse(Error, "The structure is wrong", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err = l.storage.Create(&data)
	if err != nil {
		resp := per.NewResponse(Error, "An issue occurs when try create a person", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := per.NewResponse(Message, "Person created Ok", nil)
	return c.JSON(http.StatusOK, resp)
}
