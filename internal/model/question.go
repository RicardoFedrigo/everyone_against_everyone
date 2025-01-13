package model

type Question struct {
	CommonField
	Question string `gorm:"type:text" json:"question"`
	DeckID   uint   `gorm:"not null" json:"-"`
}
