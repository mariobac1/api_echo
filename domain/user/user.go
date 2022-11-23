package user

import "github.com/mariobac1/api_/models"

type Storage interface {
	Migrate() error
	Create(*models.Login) error
	GetByEmail(Email string, Pass string) (*models.Login, bool, error)
	//Delete(ID uint) error
}
