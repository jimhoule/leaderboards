package players

import (
	"log"
	"main/players/application/services"
	"main/players/domain/factories"
	"main/players/infrastructures/persistence/fake/entities"
	"main/players/infrastructures/persistence/fake/mappers"
	"main/players/infrastructures/persistence/fake/repositories"
	"main/players/presenters/consumer/handlers"
	"main/players/presenters/http/controllers"
	"main/queue"
	"main/queue/topics"
	"main/router"
	"os"
)

func GetService() *services.PlayersService {
	return &services.PlayersService{
		PlayersFactory: &factories.PlayersFactory{},
		PlayersRepository: &repositories.FakePlayersRepository{
			PlayerEntities: []*entities.Player{},
			PlayersMapper: &mappers.PlayersMapper{},
		},
	}
}

func Init(mainRouter *router.MainRouter) {
	// Queue consumer group
	queueConsumerGroupHandler, err := queue.NewConsumerGroupHandler(
		[]string{
			os.Getenv("QUEUE_ADDRESS"),
		},
		"players_consumer_group",
	)
	if err != nil {
		log.Panic(err)
	}

	playersHandler := handlers.PlayersHandler{}

	queueConsumerGroupHandler.Handlers = map[string]queue.Handler{
		topics.PlayerScoreChanged: playersHandler.UpdateScore,
	}

	// Controller
	playersController := &controllers.PlayersController{
		PlayersService: GetService(),
	}

	mainRouter.Get("/players", playersController.GetAll)
	mainRouter.Get("/players/{id}", playersController.GetById)
	mainRouter.Post("/players", playersController.Create)
}