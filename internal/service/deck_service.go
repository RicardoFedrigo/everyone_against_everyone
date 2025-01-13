package service

import "github.com/ricardofedrigo/everyone_against_everyone/internal/model"

func ReturnDeck() []model.Deck {
	decks := []model.Deck{
		{CommonField: model.CommonField{ID: 1}, DeckName: "TESTE 1", Questions: []model.Question{{CommonField: model.CommonField{ID: 1}, Question: "abc"}}, Answers: []model.Answer{{CommonField: model.CommonField{ID: 1}, Answer: "cde"}}},
		{CommonField: model.CommonField{ID: 1}, DeckName: "TESTE 1", Questions: []model.Question{{CommonField: model.CommonField{ID: 1}, Question: "abc"}}, Answers: []model.Answer{{CommonField: model.CommonField{ID: 1}, Answer: "cde"}}},
	}

	return decks
}
