package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mariobac1/api_/models"
	"golang.org/x/crypto/bcrypt"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	MigrateUser = `CREATE TABLE IF NOT EXISTS users(
		id SERIAL NOT NULL,
		email VARCHAR(50) NOT NULL,
		password VARCHAR(350) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT users_id_pk PRIMARY KEY (id)
	)`
	CreateUser = `INSERT INTO users (email, password)
		VALUES($1, $2) RETURNING id`

	GetByEmailUser = `SELECT id, email, password, created_at, updated_at
		FROM users WHERE email = $1`
)

type user struct {
	db *sql.DB
}

func New(db *sql.DB) *user {
	return &user{db: db}
}

func (p *user) Migrate() error {
	stmt, err := p.db.Prepare(MigrateUser)
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

func (p *user) Create(m *models.Login) error {
	stmt, err := p.db.Prepare(CreateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hash, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if len(hash) == 0 {
		return errors.New("You need a password")
	}
	m.Password = string(hash)

	err = stmt.QueryRow(
		m.Email,
		m.Password,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("usera se creó correctamente")
	return nil
}

func (p *user) GetByEmail(Email string, pass string) (*models.Login, bool, error) {
	stmt, err := p.db.Prepare(GetByEmailUser)
	if err != nil {
		return &models.Login{}, false, err
	}
	defer stmt.Close()

	row, err := scanRowUser(stmt.QueryRow(Email))

	valid := isLoginValid(row, Email, pass)

	return row, valid, err
}

func isLoginValid(data *models.Login, Email string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(pass))
	return data.Email == Email && err == nil
}

func scanRowUser(s scanner) (*models.Login, error) {
	m := &models.Login{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Email,
		&m.Password,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
