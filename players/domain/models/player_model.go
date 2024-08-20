package models

type Player struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Wins   int64  `json:"wins"`
	Losses int64  `json:"losses"`
}