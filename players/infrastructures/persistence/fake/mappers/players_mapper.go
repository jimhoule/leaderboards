package mappers

import (
	"main/players/domain/models"
	"main/players/infrastructures/persistence/fake/entities"
)

type PlayersMapper struct{}

func (pm *PlayersMapper) ToModel(playerEntity *entities.Player) *models.Player {
	return &models.Player{
		Id: playerEntity.Id,
		Name: playerEntity.Name,
		Wins: playerEntity.Wins,
		Losses: playerEntity.Losses,
	}
}

func (pm *PlayersMapper) ToEntity(playerModel *models.Player) *entities.Player {
	return &entities.Player{
		Id: playerModel.Id,
		Name: playerModel.Name,
		Wins: playerModel.Wins,
		Losses: playerModel.Losses,
	}
}