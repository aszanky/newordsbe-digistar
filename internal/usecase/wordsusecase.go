package usecase

import "github.com/aszanky/newordsbe-digistar/internal/models"

func (u *usecase) AddNewWords(inp models.Word) (err error) {
	err = u.repository.AddNewWords(inp.Word, inp.Indonesian, inp.Notes)
	if err != nil {
		return
	}

	return nil
}

func (u *usecase) GetListWords() (words []models.Words, err error) {
	words, err = u.repository.GetListWord()
	if err != nil {
		return
	}

	return words, nil
}
