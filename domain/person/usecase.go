package person

import (
	"time"

	"github.com/mariobac1/api_/models"
)

var (
	//ErrIDNotFound = errors.New("This ID do not exist")
	ID uint
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(p *models.Person) error {
	return uc.storage.Create(p)
}

func (uc UseCase) GetAll() (models.Persons, error) {
	return uc.storage.GetAll()
}

func (uc UseCase) Update(m *models.Person) error {

	// if m.ID == 0 {
	// 	return ErrIDNotFound
	// }
	m.UpdatedAt = time.Now()

	return uc.storage.Update(m)
}

func (uc UseCase) GetByID(ID uint) (*models.Person, error) {
	return uc.storage.GetByID(ID)
}
