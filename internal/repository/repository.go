package repository

import (
	"github.com/aszanky/newordsbe-digistar/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AddNewWords(word, indonesian, notes string) (err error)
	GetListWord() (words []models.Words, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(
	db *sqlx.DB,
) Repository {
	return &repository{
		db: db,
	}
}
