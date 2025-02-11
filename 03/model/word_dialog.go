package model

// WordDialog - Bảng trung gian giữa Word & Dialog (Many-to-Many)
type WordDialog struct {
	DialogID string `gorm:"type:varchar(36);primaryKey;index"`
	WordID   string `gorm:"type:varchar(36);primaryKey;index"`

	Dialog Dialog `gorm:"foreignKey:DialogID;references:ID;constraint:OnDelete:CASCADE"`
	Word   Word   `gorm:"foreignKey:WordID;references:ID;constraint:OnDelete:CASCADE"`
}

func (w *WordDialog) TableName() string {
	return "intern_test.word_dialogs"
}
