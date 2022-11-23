package person

import (
	"database/sql"
	"fmt"

	"github.com/mariobac1/api_/models"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	// MigratePerson = `CREATE TABLE IF NOT EXISTS persons(
	// 	id SERIAL NOT NULL,
	// 	name VARCHAR(50) NOT NULL,
	// 	age INT NOT NULL,
	// 	communities_id INT NOT NULL,
	// 	created_at TIMESTAMP NOT NULL DEFAULT now(),
	// 	updated_at TIMESTAMP,
	// 	CONSTRAINT persons_id_pk PRIMARY KEY (id),
	// 	CONSTRAINT persons_communities_id_fk FOREIGN KEY (communities_id)
	// 	REFERENCES communities (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	// )`
	MigratePerson = `CREATE TABLE IF NOT EXISTS persons(
		id SERIAL NOT NULL,
		name VARCHAR(50) NOT NULL,
		age INT NOT NULL,
		communities_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT persons_id_pk PRIMARY KEY (id)
	)`
	CreatePerson = `INSERT INTO persons (name, age, communities_id)
		VALUES($1, $2, $3) RETURNING id`
	GetAllPerson = `SELECT id, name, age, communities_id, created_at, updated_at
		FROM persons ORDER BY id`
	GetByIDPerson = `SELECT id, name, age, communities_id, created_at, updated_at
		FROM persons WHERE id = $1`
	UpdatePerson = `UPDATE persons SET name = $1, age = $2,  communities_id = $3,
		updated_at = $4 WHERE id = $5`
	DeletePerson = `DELETE FROM persons WHERE id = $1`
)

type person struct {
	db *sql.DB
}

func New(db *sql.DB) *person {
	return &person{db: db}
}

func (p *person) Migrate() error {
	stmt, err := p.db.Prepare(MigratePerson)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de producto ejecutada correctamente")
	return nil
}

func (p *person) Create(m *models.Person) error {
	stmt, err := p.db.Prepare(CreatePerson)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		m.Age,
		m.Communities,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("Persona se creó correctamente")
	return nil
}

func (p *person) GetAll() (models.Persons, error) {
	stmt, err := p.db.Prepare(GetAllPerson)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(models.Persons, 0)
	for rows.Next() {
		m, err := scanRowPerson(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err

	}
	return ms, nil
}

// GetByID
func (p *person) GetByID(id uint) (*models.Person, error) {
	stmt, err := p.db.Prepare(GetByIDPerson)
	if err != nil {
		return &models.Person{}, err
	}
	defer stmt.Close()

	return scanRowPerson(stmt.QueryRow(id))
}

// Update
func (p *person) Update(m *models.Person) error {
	stmt, err := p.db.Prepare(UpdatePerson)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		m.Age,
		m.Communities,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el id: %d: %w", m.ID, models.ErrIDPersonDoesNotExists)
	}

	fmt.Printf("Se actualizó la persona correctamente")
	return nil

}

// Update implement the interface product.Storage
func (p *person) Delete(id uint) error {
	stmt, err := p.db.Prepare(DeletePerson)
	if err != nil {
		return err
	}
	defer stmt.Close()

	resp, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el id: %d: %w", id, models.ErrIDPersonDoesNotExists)
	}

	fmt.Println("Se eliminó la persona correctamente")
	return nil
}

func scanRowPerson(s scanner) (*models.Person, error) {
	m := &models.Person{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Age,
		&m.Communities,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

func scanRowPers(s scanner) (models.Person, error) {
	m := models.Person{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		m.ID,
		m.Name,
		m.Age,
		m.Communities,
		m.CreatedAt,
		updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
