package handler

import "github.com/kataras/iris/v12"

// GetDialogAndWords implements Handler.
func (h *handlerImpl) GetDialogAndWords(ctx iris.Context) {
	dialogID := ctx.Params().Get("dialog_id")
	if dialogID == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "dialog_id is required"})
		return
	}

	dialog, err := h.dialogRepo.GetDialogByID(ctx, dialogID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to retrieve dialog and words"})
		return
	}
	words, err := h.wordRepo.GetWordList(ctx, dialogID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to retrieve words"})
		return
	}

	ctx.JSON(iris.Map{
		"dialog": dialog,
		"words":  words,
	})
}
