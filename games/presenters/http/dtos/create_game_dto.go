package dtos

type CreateGameDto struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
}