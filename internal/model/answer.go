package model

type Answer struct {
	CommonField
	Answer string `gorm:"type:text" json:"answer"`
	DeckID uint   `gorm:"not null" json:"-"`
	Votes  int    `json:"-"`
}
