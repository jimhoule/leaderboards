package repositories

import (
	"fmt"
	"main/games/domain/models"
	"main/games/infrastructures/persistence/fake/entities"
	"main/games/infrastructures/persistence/fake/mappers"
	"time"
)

type FakeGamesRepository struct {
	GameEntities []*entities.Game
	GamesMapper *mappers.GamesMapper
	GameplayersMapper *mappers.GameplayersMapper
}

func (fgr *FakeGamesRepository) Reset() {
	fgr.GameEntities = []*entities.Game{}
}

func (fgr *FakeGamesRepository) GetPlayers(id string) (map[string]*models.Gameplayer, error) {
	// Checks if game entity exists
	gameEntity, err := getGameEntityById(fgr, id)
	if gameEntity == nil {
		return nil, err
	}
	
	return fgr.GameplayersMapper.ToModels(gameEntity.GamePlayers) , nil
}

func (fgr *FakeGamesRepository) GetById(id string) (*models.Game, error) {
	// Checks if game entity exists
	gameEntity, err := getGameEntityById(fgr, id)
	if gameEntity == nil {
		return nil, err
	}
	
	return fgr.GamesMapper.ToModel(gameEntity) , nil
}

func (fgr *FakeGamesRepository) Create(game *models.Game) (*models.Game, error) {
	fgr.GameEntities = append(
		fgr.GameEntities,
		fgr.GamesMapper.ToEntity(game),
	)

	return game, nil
}

func (fgr *FakeGamesRepository) Start(id string) (*models.Game, error) {
	// Checks if game entity exists
	gameEntity, err := getGameEntityById(fgr, id)
	if gameEntity == nil {
		return nil, err
	}

	// Checks if game has already started
	if gameEntity.StartedAt != "" {
		return nil, fmt.Errorf("game with id %s has already started", id)
	}

	// Checks if at least two players have joined the game
	if len(gameEntity.GamePlayers) < 2 {
		return nil, fmt.Errorf("game with id %s has less than 2 players", id)
	}

	gameEntity.StartedAt = time.Now().String()

	return fgr.GamesMapper.ToModel(gameEntity) , nil
}

func (fgr *FakeGamesRepository) End(id string, winnerId string) (*models.Game, error) {
	// Checks if game entity exists
	gameEntity, err := getGameEntityById(fgr, id)
	if gameEntity == nil {
		return nil, err
	}
	
	// Checks if game has not started yet
	if gameEntity.StartedAt == "" {
		return nil, fmt.Errorf("game with id %s has not started yet", id)
	}

	// Checks if game has already ended
	if gameEntity.EndedAt != "" {
		return nil, fmt.Errorf("game with id %s has already ended", id)
	}

	// Checks if winner is in the game
	_, doesExist := gameEntity.GamePlayers[winnerId]
	if !doesExist {
		return nil, fmt.Errorf("player with id %s was not in the game", winnerId)
	}

	gameEntity.GamePlayers[winnerId].IsWinner = true
	gameEntity.EndedAt = time.Now().String()		

	return fgr.GamesMapper.ToModel(gameEntity) , nil
}

func (fgr *FakeGamesRepository) Join(id string, playerId string) (*models.Game, error) {
	// Checks if game entity exists
	gameEntity, err := getGameEntityById(fgr, id)
	if gameEntity == nil {
		return nil, err
	}

	// Checks if game has already started
	if gameEntity.StartedAt != "" {
		return nil, fmt.Errorf("game with id %s has already started", id)
	}

	// Checks if game has already started
	if gameEntity.EndedAt != "" {
		return nil, fmt.Errorf("game with id %s has already ended", id)
	}

	gameEntity.GamePlayers[playerId] = &entities.Gameplayer{
		PlayerId: playerId,
		IsWinner: false,
	}

	return fgr.GamesMapper.ToModel(gameEntity) , nil
}

func getGameEntityById(fgr *FakeGamesRepository, id string) (*entities.Game, error) {
	for _, gameEntity := range fgr.GameEntities {
		if gameEntity.Id == id {
			return gameEntity, nil
		}
	}

	return nil, fmt.Errorf("game with id %s does not exist", id)
}