package models

type Game struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Mode      string `json:"mode"`
	StartedAt string `json:"startedAt"`
	EndedAt   string `json:"endedAt"`
}