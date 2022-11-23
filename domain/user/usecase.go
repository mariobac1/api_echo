package user

import (
	"errors"

	"github.com/mariobac1/api_/models"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(l *models.Login) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(l.Password), bcrypt.DefaultCost)
	if len(hash) == 0 {
		return errors.New("You need a password")
	}
	l.Password = string(hash)
	return uc.storage.Create(l)
}

func (uc UseCase) GetByEmail(Email string, pass string) (*models.Login, bool, error) {
	return uc.storage.GetByEmail(Email, pass)

}
