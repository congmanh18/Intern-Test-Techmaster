package repository

import (
	"03/model"
	"03/utils/db"
	"context"

	"gorm.io/gorm/clause"
)

type WordDialogRepo interface {
	AddWordDialog(ctx context.Context, dialogID, wordID string) error
}

type wordDialogImpl struct {
	DB *db.Database
}

// AddWordDialog implements WordDialogRepo.
func (w *wordDialogImpl) AddWordDialog(ctx context.Context, dialogID string, wordID string) error {
	wordDialog := model.WordDialog{
		DialogID: dialogID,
		WordID:   wordID,
	}

	// Thêm hoặc bỏ qua nếu đã tồn tại
	err := w.DB.Executor.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "dialog_id"}, {Name: "word_id"}},
			DoNothing: true,
		}).Create(&wordDialog).Error

	return err
}

func NewWordDialogRepo(db *db.Database) WordDialogRepo {
	return &wordDialogImpl{DB: db}
}
