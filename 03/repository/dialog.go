package repository

import (
	"03/model"
	"03/utils/db"
	"context"
)

type DialogRepo interface {
	AddDialog(ctx context.Context, content string) (string, error)
	GetDialogByID(ctx context.Context, dialogID string) (model.Dialog, error)
	GetAllIDDialogs(ctx context.Context) ([]string, error)
}

type dialogImpl struct {
	DB *db.Database
}

func (d *dialogImpl) GetAllIDDialogs(ctx context.Context) ([]string, error) {
	var dialogs []string

	if err := d.DB.Executor.WithContext(ctx).
		Model(&model.Dialog{}).
		Select("id").
		Order("id DESC").
		Find(&dialogs).Error; err != nil {
		return nil, err
	}
	return dialogs, nil
}

func (d *dialogImpl) GetDialogByID(ctx context.Context, dialogID string) (model.Dialog, error) {
	var dialog model.Dialog
	err := d.DB.Executor.WithContext(ctx).
		Model(&model.Dialog{}).
		Where("id = ?", dialogID).
		First(&dialog).Error

	return dialog, err
}

func (d *dialogImpl) AddDialog(ctx context.Context, content string) (string, error) {
	dialog := model.Dialog{
		ID:      newULID(),
		Lang:    model.Vietnamese,
		Content: content,
	}
	if err := d.DB.Executor.WithContext(ctx).
		Create(&dialog).Error; err != nil {
		return dialog.ID, err
	}
	return dialog.ID, nil
}

func NewDialogRepo(db *db.Database) DialogRepo {
	return &dialogImpl{DB: db}
}
