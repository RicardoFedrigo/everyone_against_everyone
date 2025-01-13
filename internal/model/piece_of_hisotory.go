package model

type PieceOfHisotory struct {
	CommonField
	Sequence int    `json:"sequence"`
	Phrase   string `json:"phrase"`
}
