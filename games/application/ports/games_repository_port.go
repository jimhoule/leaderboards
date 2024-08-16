package ports

import "main/games/domain/models"

type GameRepositoryPort interface {
	GetById(id string) (*models.Game, error)
	Create(game *models.Game) (*models.Game, error)
	Start(id string) (*models.Game, error)
	End(id string, winnerId string) (*models.Game, error)
	Join(id string, playerId string) (*models.Game, error)
}