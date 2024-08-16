package mappers

import (
	"main/games/domain/models"
	"main/games/infrastructures/persistence/fake/entities"
)

type GamesMapper struct{}

func (gm *GamesMapper) ToModel(gameEntity *entities.Game) *models.Game {
	return &models.Game{
		Id: gameEntity.Id,
		Name: gameEntity.Name,
		Mode: gameEntity.Mode,
		StartedAt: gameEntity.StartedAt,
		EndedAt: gameEntity.EndedAt,
	}
}

func (gm *GamesMapper) ToEntity(gameModel *models.Game) *entities.Game {
	return &entities.Game{
		Id: gameModel.Id,
		Name: gameModel.Name,
		Mode: gameModel.Mode,
		StartedAt: gameModel.StartedAt,
		EndedAt: gameModel.EndedAt,
		GamePlayers: map[string]*entities.Gameplayer{},
	}
}