package handler

import (
	"03/repository"

	"github.com/kataras/iris/v12"
)

type Handler interface {
	AutoProcessHandler(ctx iris.Context)
}

type handlerImpl struct {
	dialogRepo     repository.DialogRepo
	wordRepo       repository.WordRepo
	wordDialogRepo repository.WordDialogRepo
}

type HandlerInject struct {
	DialogRepo     repository.DialogRepo
	WordRepo       repository.WordRepo
	WordDialogRepo repository.WordDialogRepo
}

func NewHandler(inj HandlerInject) Handler {
	return &handlerImpl{
		dialogRepo:     inj.DialogRepo,
		wordRepo:       inj.WordRepo,
		wordDialogRepo: inj.WordDialogRepo,
	}
}
