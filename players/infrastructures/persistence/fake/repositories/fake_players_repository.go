package repositories

import (
	"fmt"
	"main/players/domain/models"
	"main/players/infrastructures/persistence/fake/entities"
	"main/players/infrastructures/persistence/fake/mappers"
)

type FakePlayersRepository struct {
	PlayerEntities []*entities.Player
	PlayersMapper *mappers.PlayersMapper
}

func (fpr *FakePlayersRepository) Reset() {
	fpr.PlayerEntities = []*entities.Player{}
}

func (fpr *FakePlayersRepository) GetAll() ([]*models.Player, error) {
	players := []*models.Player{}

	for _, playerEntity := range fpr.PlayerEntities {
		players = append(players, fpr.PlayersMapper.ToModel(playerEntity))
	}

	return players, nil
}

func (fpr *FakePlayersRepository) GetById(id string) (*models.Player, error) {
	for _, playerEntity := range fpr.PlayerEntities {
		if playerEntity.Id == id {
			return fpr.PlayersMapper.ToModel(playerEntity), nil
		}
	}

	return nil, fmt.Errorf("player with id %s does not exist", id)
}

func (fpr *FakePlayersRepository) Create(player *models.Player) (*models.Player, error) {
	fpr.PlayerEntities = append(
		fpr.PlayerEntities,
		fpr.PlayersMapper.ToEntity(player),
	)

	return player, nil
}

func (fpr *FakePlayersRepository) UpdateScore(id string, isWinner bool) (*models.Player, error) {
	player, err := fpr.GetById(id)
	if err != nil {
		return nil, err
	}

	if isWinner {
		player.Wins++
	} else {
		player.Losses++
	}

	return nil, nil
}