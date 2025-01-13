package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/ricardofedrigo/everyone_against_everyone/internal/model"
)

type GameLogicService interface {
	// StartGame starts the game
	StartGame(room model.Room) bool
	// EndGame ends the game
	EndGame() error
	// NextQuestion returns the next question
	NextQuestion() error
	// AnswerQuestion answers the question
	AnswerQuestion() error
	// GetQuestion returns the current question
	GetQuestion() model.Question
	// SetAnswer sets the answer with more votes
	SetAnswer() error
	// GetPeiceOfHistory returns the piece of history
	GetPeiceOfHistory() model.PieceOfHisotory
	// PieceOfHistory returns the set arry piece of history
	initPieceOfHistory() []model.PieceOfHisotory
	AddPlayer(player model.Player) error
}

// GameLogicServiceImpl is the implementation of the GameLogic interface
type GameLogicServiceImpl struct {
	room            model.Room
	history         model.History
	pieceOfHisotory []model.PieceOfHisotory
	question        []model.Question
	round           int
	rng             *rand.Rand
}

func (g *GameLogicServiceImpl) initPieceOfHistory() *[]model.PieceOfHisotory {
	return &[]model.PieceOfHisotory{
		{
			Sequence: 0,
			Phrase:   "O Brasil é o país do futebol",
		},
		{
			Sequence: 0,
			Phrase:   "O Brasil é o país do carnaval",
		},
		{
			Sequence: 0,
			Phrase:   "O Brasil é o país da corrupção",
		},
		{
			Sequence: 0,
			Phrase:   "O Brasil é o país da impunidade",
		},
		{
			Sequence: 0,
			Phrase:   "O Brasil é o país da desigualdade",
		},
	}
}

func (g *GameLogicServiceImpl) AddPlayer(player model.Player) error {
	g.room.Players = append(g.room.Players, player)
	return nil
}

// NewGameLogicImpl returns a new GameLogicImpl
func (g *GameLogicServiceImpl) NewGameLogicImpl() *GameLogicServiceImpl {
	return &GameLogicServiceImpl{
		room: model.Room{},
		rng:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// StartGame starts the game
func (g *GameLogicServiceImpl) StartGame(room model.Room) bool {
	for _, player := range room.Players {
		if !player.IsReady {
			return false
		}
	}

	return true
}

func (g *GameLogicServiceImpl) GetQuestion() model.Question {
	return model.Question{}
}

func (g *GameLogicServiceImpl) NextQuestion() error {
	return nil
}

func (g *GameLogicServiceImpl) GetPeiceOfHistory() model.PieceOfHisotory {

	index := g.rng.Intn(len(g.history.PieceOfHisotory))

	pieceOfHisotory := g.pieceOfHisotory[index]
	g.history.PieceOfHisotory = append(g.history.PieceOfHisotory[:index], g.history.PieceOfHisotory[index+1:]...)
	return pieceOfHisotory
}

func (g *GameLogicServiceImpl) SetAnswer(answers **[]model.Answer) error {
	answerWithMoreVotes := model.Answer{}

	if answers == nil {
		return errors.New("must have one answer")
	}

	for _, answer := range **answers {
		if answer.Votes > answerWithMoreVotes.Votes {
			answerWithMoreVotes = answer
		}
	}

	g.history.PieceOfHisotory = append(g.history.PieceOfHisotory, model.PieceOfHisotory{
		Sequence: g.round,
		Phrase:   answerWithMoreVotes.Answer,
	})

	g.round++

	return nil
}
