package ports

import "main/players/domain/models"

type PlayersRepositoryPort interface {
	GetAll() ([]*models.Player, error)
	GetById(id string) (*models.Player, error)
	Create(player *models.Player) (*models.Player, error)
	UpdateScore(id string, isWinner bool) (*models.Player, error)
}