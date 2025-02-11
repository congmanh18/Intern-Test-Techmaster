package server

import (
	"03/handler"

	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application, h handler.Handler) {
	api := app.Party("/api")

	api.Get("/dialogs/ids", h.GetAllIDDialogs)           // Route để lấy tất cả ID của dialogs
	api.Get("/dialogs/{dialog_id}", h.GetDialogAndWords) // Route để lấy hội thoại và danh sách từ
}
