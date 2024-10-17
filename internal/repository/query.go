package repository

const (
	QueryAddNewWords = `INSERT INTO translation(word, indonesian, notes) VALUES ($1, $2, $3)`
)
