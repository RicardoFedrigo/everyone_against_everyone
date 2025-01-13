package model

type Deck struct {
	CommonField
	DeckName  string     `gorm:"varchar(255)" json:deckName`
	Questions []Question `gorm:"foreignKey:DeckID" json:questions`
	Answers   []Answer   `gorm:"foreignKey:DeckID" json:answers`
}
