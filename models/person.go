package models

import "time"

type Person struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Age         uint8     `json:"age"`
	Communities uint      `json:"communities"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Persons []*Person
