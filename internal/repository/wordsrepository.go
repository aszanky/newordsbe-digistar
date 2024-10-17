package repository

import "github.com/aszanky/newordsbe-digistar/internal/models"

func (r *repository) AddNewWords(word, indonesia, notes string) (err error) {
	queryAddNewWords := `INSERT INTO translation(word, indonesian, notes) VALUES ($1, $2, $3)`

	_, err = r.db.Exec(queryAddNewWords, word, indonesia, notes)
	if err != nil {
		return
	}

	return nil
}

func (r *repository) GetListWord() (words []models.Words, err error) {
	queryGetWords := `SELECT * FROM translation`

	err = r.db.Select(&words, queryGetWords)
	if err != nil {
		return
	}

	return words, nil
}
