package games

import (
	"main/games/application/services"
	"main/games/domain/factories"
	"main/games/domain/models"
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
			Games: []*models.Game{},
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
}