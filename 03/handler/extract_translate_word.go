package handler

import (
	"03/model"
	"encoding/json"
	"errors"
	"regexp"

	"github.com/kataras/iris/v12"
)

func extractTranslatedWords(response string) (string, error) {
	// Regex để tìm JSON chứa "translated_words"
	re := regexp.MustCompile(`\{[\s\S]*?"translated_words":\s*\[[\s\S]*?\]\}`)
	match := re.FindString(response)

	if match == "" {
		return "", errors.New("không tìm thấy phần translated_words trong phản hồi API")
	}

	// Chuẩn hóa JSON (loại bỏ khoảng trắng dư thừa)
	var formatted map[string]interface{}
	if err := json.Unmarshal([]byte(match), &formatted); err != nil {
		return "", errors.New("lỗi phân tích JSON từ phản hồi API")
	}

	// Chuyển đổi lại thành JSON đẹp có xuống dòng
	formattedJSON, err := json.MarshalIndent(formatted, "", "  ")
	if err != nil {
		return "", errors.New("lỗi định dạng lại JSON")
	}

	return string(formattedJSON), nil
}

func (h *handlerImpl) SaveTranslatedWords(ctx iris.Context, dialogID, translatedWords string) error {
	var parsedData struct {
		TranslatedWords []model.TranslatedWord `json:"translated_words"`
	}

	err := json.Unmarshal([]byte(translatedWords), &parsedData)
	if err != nil {
		return err
	}

	for _, word := range parsedData.TranslatedWords {
		wordID, err := h.wordRepo.AddWord(ctx, word)
		if err != nil {
			return err
		}

		// Liên kết từ với hội thoại trong bảng word_dialog
		err = h.wordDialogRepo.AddWordDialog(ctx, dialogID, wordID)
		if err != nil {
			return err
		}
	}
	return nil
}
