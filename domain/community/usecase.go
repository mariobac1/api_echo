package community

import "github.com/mariobac1/api_/models"

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(p *models.Community) error {
	return uc.storage.Create(p)
}

// func (uc UseCase) GetAll() (models.Communities, error) {
// 	return uc.storage.GetAll()
// }

// func (uc UseCase) GetByID(ID uint) (models.Community, error) {
// 	return uc.storage.GetByID(ID)
// }
