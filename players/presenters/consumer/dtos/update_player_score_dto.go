package dtos

type UpdatePlayerScoreDto struct {
	Id       string `json:"id"`
	IsWinner bool   `json:"isWnner"`
}