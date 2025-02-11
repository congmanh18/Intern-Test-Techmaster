package repository

import (
	"03/model"
	"03/utils/db"
	"context"
)

type WordRepo interface {
	AddWord(ctx context.Context, word model.TranslatedWord) (string, error)
	DeleteWord(ctx context.Context, word model.Word) (bool, error)
	GetWordList(ctx context.Context, limit, offset int) ([]model.Word, error)
}

type wordImpl struct {
	DB *db.Database
}

// AddWord implements WordRepo.
func (w *wordImpl) AddWord(ctx context.Context, word model.TranslatedWord) (string, error) {
	newWord := model.Word{
		ID:        newULID(),
		Lang:      model.Vietnamese,
		Content:   word.Vi,
		Translate: word.En,
	}
	if err := w.DB.Executor.Create(&newWord).Error; err != nil {
		return newWord.ID, err
	}
	return newWord.ID, nil
}

// DeleteWord implements WordRepo.
func (w *wordImpl) DeleteWord(ctx context.Context, word model.Word) (bool, error) {
	panic("unimplemented")
}

// GetWordList implements WordRepo.
func (w *wordImpl) GetWordList(ctx context.Context, limit int, offset int) ([]model.Word, error) {
	panic("unimplemented")
}

func NewWordRepo(db *db.Database) WordRepo {
	return &wordImpl{DB: db}
}
