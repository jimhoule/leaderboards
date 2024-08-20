package services

import (
	"main/players/application/payloads"
	"main/players/application/ports"
	"main/players/domain/factories"
	"main/players/domain/models"
)

type PlayersService struct {
	PlayersFactory *factories.PlayersFactory
	PlayersRepository ports.PlayersRepositoryPort
}

func (ps *PlayersService) GetAll() ([]*models.Player, error) {
	return ps.PlayersRepository.GetAll()
}

func (ps *PlayersService) GetById(id string) (*models.Player, error) {
	return ps.PlayersRepository.GetById(id)
}

func (ps *PlayersService) Create(createPlayerPayload *payloads.CreatePlayerPayload) (*models.Player, error) {
	player := ps.PlayersFactory.New(createPlayerPayload.Name)

	return ps.PlayersRepository.Create(player)
}

func (ps *PlayersService) UpdateScore(updatePlayerScorePayload *payloads.UpdatePlayerScorePayload) (*models.Player, error) {
	return ps.PlayersRepository.UpdateScore(updatePlayerScorePayload.Id, updatePlayerScorePayload.IsWinner)
}