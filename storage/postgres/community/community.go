package community

import (
	"database/sql"
	"fmt"

	"github.com/mariobac1/api_/models"
)

const (
	MigrateCommunity = ` CREATE TABLE IF NOT EXISTS communities(
		id Serial NOT NULL,
		name VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT communities_id_pk PRIMARY KEY (id)
		)`
	CreateCommunity = `INSERT INTO communities (name)
		VALUES($1) RETURNING id`
	GetAllCommunity = `SELECT id, name, created_at, updated_at
		FROM communities`
	GetByIDCommunity = GetAllCommunity + `WHERE id = $1`
	UpdateCommunity  = `UPDATE communities SET name = $1, updated_at = $2, WHERE id = $3`
)

type community struct {
	db *sql.DB
}

func New(db *sql.DB) *community {
	return &community{db: db}
}

func (p *community) Migrate() error {
	stmt, err := p.db.Prepare(MigrateCommunity)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de Community ejecutada correctamente")
	return nil
}

func (p *community) Create(c *models.Community) error {
	stmt, err := p.db.Prepare(CreateCommunity)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		c.Name,
	).Scan(&c.ID)
	if err != nil {
		return err
	}

	fmt.Println("Se creó la comunidad correctamente")
	return nil
}
