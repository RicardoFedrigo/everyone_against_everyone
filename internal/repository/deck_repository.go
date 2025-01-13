package repository

import (
	"github.com/ricardofedrigo/everyone_against_everyone/internal/model"
	"gorm.io/gorm"
)

type DeckRepositoryInterface interface {
	findAll() ([]model.Deck, error)
}

type DeckRepository struct {
	DB *gorm.DB
}

func (r *DeckRepository) findAll() ([]model.Deck, error) {
	var decks []model.Deck

	err := r.DB.Joins("Questions").Joins("Answer").Find(&decks).Error

	if err != nil {
		return nil, err
	}

	return decks, nil
}
