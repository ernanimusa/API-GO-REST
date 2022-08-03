package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
	Obs      string `json:"observacao"`
}

var Personalidades []Personalidade
