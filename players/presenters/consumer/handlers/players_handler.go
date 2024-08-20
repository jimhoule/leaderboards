package handlers

import (
	"encoding/json"
	"main/players/application/payloads"
	"main/players/application/services"
	"main/players/presenters/consumer/dtos"
)

type PlayersHandler struct{
	PlayersService *services.PlayersService
}

func (ph *PlayersHandler) UpdateScore(body []byte) error {
	var updatePlayerScoreDto dtos.UpdatePlayerScoreDto
	err := json.Unmarshal(body, &updatePlayerScoreDto)
	if err != nil {
		return err
	}

	_, err = ph.PlayersService.UpdateScore(&payloads.UpdatePlayerScorePayload{
		Id: updatePlayerScoreDto.Id,
		IsWinner: updatePlayerScoreDto.IsWinner,
	})
	if err != nil {
		return err
	}

	return nil
}