package handler

import (
	"03/model"
	"03/utils/groq"

	"github.com/kataras/iris/v12"
	"github.com/russross/blackfriday/v2"
)

type Promt string

const (
	GenSixSentence Promt = `Tạo một hội thoại bằng tiếng Việt, gồm 6 câu, ngắn gọn, đơn giản, hỏi đường đi đến hồ Hoàn Kiếm ở Hà nội giữa một bạn James và tên Lan. Chỉ xuất ra hội thoại không cần giải thích.`
	FilterOut      Promt = `Từ hội thoại trên hãy lọc ra danh sách các từ quan trọng, bỏ qua danh từ tên riêng cần học. Không cần giải thích xuất kết quả ra dạng JSON trong thẻ "words" Ví dụ {"word":[]}.`
	Translate      Promt = `Dịch từng từ trong danh sách dưới sang tiếng Anh rồi tra ̉ JSON gồm mảng trong đó mỗi phần tử sẽ gồm từ tiếmh Việt và từ tiếng Anh tương đương. Không cần giải thích. Ví dụ {"translated_words": [{"vi":"bạn", "en":"you"}]}`
)

// AutoProcessHandler thực hiện tự động 3 bước
func (h *handlerImpl) AutoProcessHandler(ctx iris.Context) {
	var req model.Request
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Lỗi khi đọc dữ liệu"})
		return
	}

	apiURL := "https://api.groq.com/openai/v1/chat/completions"
	apiKey := "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"

	// Bước 1: Tạo đoạn hội thoại 6 câu
	dialogue, err := groq.CallGroqAPI(apiURL, apiKey, string(GenSixSentence), req.Model)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi khi gọi API bước 1", "details": err.Error()})
		return
	}
	htmlDialogue := string(blackfriday.Run([]byte(dialogue)))
	dialogID, err := h.dialogRepo.AddDialog(ctx, htmlDialogue)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi khi lưu đoạn hội thoại", "details": err.Error()})
		return
	}

	// Bước 2: Lọc danh sách từ quan trọng
	filterPrompt := string(FilterOut) + "\n\n" + dialogue
	importantWords, err := groq.CallGroqAPI(apiURL, apiKey, filterPrompt, req.Model)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi khi gọi API bước 2", "details": err.Error()})
		return
	}
	htmlImportantWords := string(blackfriday.Run([]byte(importantWords)))

	// Bước 3: Dịch danh sách từ sang tiếng Anh
	translatePrompt := string(Translate) + "\n\n" + importantWords
	translatedWords, err := groq.CallGroqAPI(apiURL, apiKey, translatePrompt, req.Model)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi khi gọi API bước 3", "details": err.Error()})
		return
	}
	htmlTranslatedWords := string(blackfriday.Run([]byte(translatedWords)))
	cleanedJSON, err := extractTranslatedWords(translatedWords)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi khi lọc dữ liệu JSON", "details": err.Error()})
		return
	}

	err = h.SaveTranslatedWords(ctx, dialogID, cleanedJSON)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi khi lưu dữ liệu", "details": err.Error()})
		return
	}
	// Trả về JSON kết quả
	ctx.JSON(model.ResponseData{
		Dialogue:        htmlDialogue,
		ImportantWords:  htmlImportantWords,
		TranslatedWords: htmlTranslatedWords,
	})
}
