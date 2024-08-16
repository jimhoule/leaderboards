package services

import (
	"main/games/application/payloads"
	"main/games/application/ports"
	"main/games/domain/factories"
	"main/games/domain/models"
)

type GamesService struct {
	GamesFactory *factories.GamesFactory
	GamesRepository ports.GameRepositoryPort
}

func (gs *GamesService) GetById(id string) (*models.Game, error) {
	return gs.GamesRepository.GetById(id)
}

func (gs *GamesService) Create(createGamePayload *payloads.CreateGamePayload) (*models.Game, error) {
	game := gs.GamesFactory.New(createGamePayload.Name, createGamePayload.Mode)

	return gs.GamesRepository.Create(game)
}

func (gs *GamesService) Start(startGamePayload *payloads.StartGamePayload) (*models.Game, error) {
	return gs.GamesRepository.Start(startGamePayload.Id)
}

func (gs *GamesService) End(endGamePayload *payloads.EndGamePayload) (*models.Game, error) {
	return gs.GamesRepository.End(endGamePayload.Id, endGamePayload.WinnerId)
}

func (gs *GamesService) Join(joinGamePayload *payloads.JoinGamePayload) (*models.Game, error) {
	return gs.GamesRepository.Join(joinGamePayload.Id, joinGamePayload.PlayerId)
}