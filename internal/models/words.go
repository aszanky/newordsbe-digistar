package models

import (
	"database/sql"
	"time"
)

type Word struct {
	Word       string `json:"word"`
	Indonesian string `json:"indonesian"`
	Notes      string `json:"notes"`
}

type Words struct {
	ID         int64        `json:"id" db:"id"`
	Word       string       `json:"word" db:"word"`
	Indonesian string       `json:"indonesian" db:"indonesian"`
	Notes      string       `json:"notes" db:"notes"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at" db:"updated_at"`
}
