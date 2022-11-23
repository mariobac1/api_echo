package community

import "github.com/mariobac1/api_/models"

//******************
//**** PORT OUT ****
//******************

type Storage interface {
	Migrate() error
	Create(*models.Community) error
	//Update(*models.Community) error
	//GetAll() (models.Communities, error)
	//GetByID(ID uint) (models.Community, error)
	//Delete(ID uint) error
}

//******************
//**** PORT IN ****
//******************
