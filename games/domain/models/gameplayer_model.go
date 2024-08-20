package models

type Gameplayer struct {
	PlayerId string `json:"playerId"`
	IsWinner bool   `json:"isWinner"`
}