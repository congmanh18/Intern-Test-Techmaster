package model

type Language string

const (
	Vietnamese Language = "vi"
	English    Language = "en"
)

// Word - Bảng lưu trữ từ vựng và bản dịch
type Word struct {
	ID        string   `gorm:"type:varchar(36);primaryKey"`
	Lang      Language `gorm:"type:varchar(2);not null" json:"lang"` // "vi" hoặc "en"
	Content   string   `gorm:"type:text;not null" json:"content"`    // Từ gốc
	Translate string   `gorm:"type:text;not null" json:"translate"`  // Bản dịch

	Dialogs []*Dialog `gorm:"many2many:word_dialogs;" json:"dialogs"`
}

type TranslatedWord struct {
	Vi string `json:"vi"`
	En string `json:"en"`
}

func (w *Word) TableName() string {
	return "intern_test.words"
}
