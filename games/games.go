package games

import (
	"log"
	"main/games/application/services"
	"main/games/domain/factories"
	"main/games/infrastructures/persistence/fake/entities"
	"main/games/infrastructures/persistence/fake/mappers"
	"main/games/infrastructures/persistence/fake/repositories"
	"main/games/presenters/http/controllers"
	"main/queue"
	"main/router"
	"main/uuid"
	"os"
)

func GetService() *services.GamesService {
	return &services.GamesService{
		GamesFactory: &factories.GamesFactory{
			UuidService: uuid.GetService(),
		},
		GamesRepository: &repositories.FakeGamesRepository{
			GameEntities: []*entities.Game{},
			GamesMapper: &mappers.GamesMapper{},
			GameplayersMapper: &mappers.GameplayersMapper{},
		},
	}
}

func Init(mainRouter *router.MainRouter) {
	queueProducerHandler, err := queue.NewProducerHandler([]string{
		os.Getenv("QUEUE_ADDRESS"),
	})
	if err != nil {
		log.Panic(err)
	}

	gamesController := &controllers.GamesController{
		GamesService: GetService(),
		QueueProducerHandler: &queueProducerHandler,
	}

	mainRouter.Get("/games/{id}", gamesController.GetById)
	mainRouter.Post("/games", gamesController.Create)
	mainRouter.Patch("/games/{id}/start", gamesController.Start)
	mainRouter.Patch("/games/{id}/end", gamesController.End)
	mainRouter.Patch("/games/{id}/join", gamesController.Join)
}