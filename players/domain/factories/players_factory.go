package factories

import (
	"main/players/domain/models"
	"main/uuid/services"
)

type PlayersFactory struct {
	UuidService services.UuidService
}

func (pf *PlayersFactory) New(name string) *models.Player {
	return &models.Player{
		Id: pf.UuidService.Generate(),
		Name: name,
		Wins: 0,
		Losses: 0,
	}
}