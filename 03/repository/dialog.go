package repository

import (
	"03/model"
	"03/utils/db"
	"context"
)

type DialogRepo interface {
	AddDialog(ctx context.Context, content string) (string, error)
	// GetDialogs() ([]model.Dialog, error)
}

type dialogImpl struct {
	DB *db.Database
}

func (d *dialogImpl) AddDialog(ctx context.Context, content string) (string, error) {
	dialog := model.Dialog{
		ID:      newULID(),
		Lang:    model.Vietnamese,
		Content: content,
	}
	if err := d.DB.Executor.Create(&dialog).Error; err != nil {
		return dialog.ID, err
	}
	return dialog.ID, nil
}

// func (d *dialogImpl) GetDialogs() ([]model.Dialog, error) {
// 	panic("unimplemented")
// }

func NewDialogRepo(db *db.Database) DialogRepo {
	return &dialogImpl{DB: db}
}
