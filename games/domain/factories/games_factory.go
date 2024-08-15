package factories

import (
	"main/games/domain/models"
	"main/uuid/services"
)

type GamesFactory struct {
	UuidService services.UuidService
}

func (gf *GamesFactory) New(name string, mode string) *models.Game {
	return &models.Game{
		Id: gf.UuidService.Generate(),
		Name: name,
		Mode: mode,
		StartedAt: "",
		EndedAt: "",
	}
}