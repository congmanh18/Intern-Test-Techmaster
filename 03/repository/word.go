package repository

import (
	"03/model"
	"03/utils/db"
	"context"
	"fmt"
)

type WordRepo interface {
	AddWord(ctx context.Context, word model.TranslatedWord) (string, error)
	GetWordList(ctx context.Context, dialogID string) ([]*model.Word, error)
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
	if err := w.DB.Executor.WithContext(ctx).
		Create(&newWord).Error; err != nil {
		return newWord.ID, err
	}
	return newWord.ID, nil
}

func (w *wordImpl) GetWordList(ctx context.Context, dialogID string) ([]*model.Word, error) {
	var wordIDs []string

	// Lấy tất cả các word_id liên quan đến dialogID từ bảng word_dialogs
	if err := w.DB.Executor.WithContext(ctx).
		Table("intern_test.word_dialogs").
		Select("word_id").
		Where("dialog_id = ?", dialogID).
		Pluck("word_id", &wordIDs).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve word IDs by dialog ID: %w", err)
	}

	if len(wordIDs) == 0 {
		return []*model.Word{}, nil // Trả về danh sách rỗng nếu không có từ nào
	}

	var words []*model.Word

	// Lấy tất cả các Word từ bảng words dựa trên danh sách word_id
	if err := w.DB.Executor.WithContext(ctx).
		Table("intern_test.words").
		Where("id IN (?)", wordIDs).
		Find(&words).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve words: %w", err)
	}

	return words, nil
}

func NewWordRepo(db *db.Database) WordRepo {
	return &wordImpl{DB: db}
}
