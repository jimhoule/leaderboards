package mappers

import (
	"main/games/domain/models"
	"main/games/infrastructures/persistence/fake/entities"
)

type GameplayersMapper struct{}

func (gm *GameplayersMapper) ToModels(gameplayerEntitiesMap map[string]*entities.Gameplayer) map[string]*models.Gameplayer {
	gamePlayersMap := map[string]*models.Gameplayer{}

	for key, gameplayerEntity := range gameplayerEntitiesMap {
		gamePlayer := &models.Gameplayer{
			PlayerId: gameplayerEntity.PlayerId,
			IsWinner: gameplayerEntity.IsWinner,
		}

		gamePlayersMap[key] = gamePlayer
	}

	return gamePlayersMap
}