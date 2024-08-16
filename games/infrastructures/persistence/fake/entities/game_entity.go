package entities

type Game struct {
	Id          string
	Name        string
	Mode        string
	StartedAt   string
	EndedAt     string
	GamePlayers map[string]*Gameplayer
}