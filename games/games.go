package games

import (
	"main/games/application/services"
	"main/games/domain/factories"
	"main/games/infrastructures/persistence/fake/entities"
	"main/games/infrastructures/persistence/fake/mappers"
	"main/games/infrastructures/persistence/fake/repositories"
	"main/games/presenters/http/controllers"
	"main/router"
	"main/uuid"
)

func GetService() *services.GamesService {
	return &services.GamesService{
		GamesFactory: &factories.GamesFactory{
			UuidService: uuid.GetService(),
		},
		GamesRepository: &repositories.FakeGamesRepository{
			GameEntities: []*entities.Game{},
			GamesMapper: &mappers.GamesMapper{},
		},
	}
}

func Init(mainRouter *router.MainRouter) {
	gamesController := &controllers.GamesController{
		GamesService: GetService(),
	}

	mainRouter.Get("/games/{id}", gamesController.GetById)
	mainRouter.Post("/games", gamesController.Create)
	mainRouter.Patch("/games/{id}/start", gamesController.Start)
	mainRouter.Patch("/games/{id}/end", gamesController.End)
	mainRouter.Patch("/games/{id}/join", gamesController.Join)
}