package handler

import "github.com/kataras/iris/v12"

// GetAllIDDialogs implements Handler.
func (h *handlerImpl) GetAllIDDialogs(ctx iris.Context) {
	ids, err := h.dialogRepo.GetAllIDDialogs(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to get dialog IDs"})
		return
	}

	ctx.JSON(iris.Map{"dialog_ids": ids})
}
