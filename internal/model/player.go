package model

type Player struct {
	Name            string
	PieceOfHisotory []PieceOfHisotory
	Answers         []Answer
	Points          int
	IsReady         bool
}
