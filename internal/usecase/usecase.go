package usecase

import (
	"github.com/aszanky/newordsbe-digistar/internal/models"
	"github.com/aszanky/newordsbe-digistar/internal/repository"
)

type Usecase interface {
	AddNewWords(inp models.Word) (err error)
	GetListWords() (words []models.Words, err error)
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(
	rep repository.Repository,
) Usecase {
	return &usecase{
		repository: rep,
	}
}
