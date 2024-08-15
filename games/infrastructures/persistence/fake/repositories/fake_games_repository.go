package repositories

import (
	"fmt"
	"main/games/domain/models"
	"time"
)

type FakeGamesRepository struct {
	Games []*models.Game
}

func (fgr *FakeGamesRepository) Reset() {
	fgr.Games = []*models.Game{}
}

func (fgr *FakeGamesRepository) GetById(id string) (*models.Game, error) {
	for _, game := range fgr.Games {
		if game.Id == id {
			return game , nil
		}
	}

	return nil, fmt.Errorf("game with id %s does not exist", id)
}

func (fgr *FakeGamesRepository) Create(game *models.Game) (*models.Game, error) {
	fgr.Games = append(fgr.Games, game)

	return game, nil
}

func (fgr *FakeGamesRepository) Start(id string) (*models.Game, error) {
	game, err := fgr.GetById(id)
	if err != nil {
		return nil, err
	}

	game.StartedAt = time.Now().String()

	return game, nil
}

func (fgr *FakeGamesRepository) End(id string) (*models.Game, error) {
	game, err := fgr.GetById(id)
	if err != nil {
		return nil, err
	}

	game.EndedAt = time.Now().String()

	return game, nil
}