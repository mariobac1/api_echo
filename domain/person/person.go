package person

import "github.com/mariobac1/api_/models"

//******************
//**** PORT OUT ****
//******************

type Storage interface {
	Migrate() error
	Create(*models.Person) error
	Update(*models.Person) error
	GetAll() (models.Persons, error)
	GetByID(ID uint) (*models.Person, error)
	Delete(ID uint) error
}

//******************
//**** PORT IN ****
//******************

type Person interface {
	Migrate() error
	Create(*models.Person) error
	Update(*models.Person) error
	GetAll() (models.Persons, error)
	GetByID(ID uint) (*models.Person, error)
	Delete(ID uint) error
}
