package service

import "github.com/ricardofedrigo/everyone_against_everyone/internal/model"

type PlayerService interface {
	GetAllAnswer() *model.Answer
	GetAnswersById(id uint) *[]model.Answer
}

// PlayerServiceImpl is the implementation of the PlayerService interface	'
type PlayerServiceImpl struct {
	player model.Player
}

// new player
func NewPlayer(name string, answer []model.Answer) *PlayerServiceImpl {
	return &PlayerServiceImpl{
		player: model.Player{
			Name:    name,
			Answers: answer,
		},
	}
}

func (p *PlayerServiceImpl) GetAnswersById(id uint) *model.Answer {

	for _, answer := range p.player.Answers {
		if answer.ID == id {
			return &answer
		}
	}
	return nil
}

func (p *PlayerServiceImpl) GetAllAnswer() *[]model.Answer {
	return &p.player.Answers
}
