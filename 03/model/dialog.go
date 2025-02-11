package model

// Dialog - Bảng lưu trữ hội thoại
type Dialog struct {
	ID      string   `gorm:"type:varchar(36);primaryKey"`
	Lang    Language `gorm:"type:varchar(2);not null" json:"lang"` // "vi" hoặc "en"
	Content string   `gorm:"type:text;not null" json:"content"`    // Nội dung hội thoại

	Words []*Word `gorm:"many2many:word_dialogs;" json:"words"`
}

func (d *Dialog) TableName() string {
	return "intern_test.dialogs"
}
