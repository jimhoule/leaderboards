package controllers

import (
	"main/games/application/payloads"
	"main/games/application/services"
	"main/games/presenters/http/dtos"
	"main/router"
	"main/router/utils"
	"net/http"
)

type GamesController struct{
	GamesService *services.GamesService
}

func (gc *GamesController) GetById(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	game, err := gc.GamesService.GetById(id)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, game)
}

func (gc *GamesController) Create(writer http.ResponseWriter, request *http.Request) {
	var createGameDto dtos.CreateGameDto
	err := utils.ReadHttpRequestBody(writer, request, &createGameDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	game, err := gc.GamesService.Create(&payloads.CreateGamePayload{
		Name: createGameDto.Name,
		Mode: createGameDto.Mode,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusCreated, game)
}

func (gc *GamesController) Start(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	game, err := gc.GamesService.Start(&payloads.StartGamePayload{
		Id: id,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, game)
}

func (gc *GamesController) End(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	game, err := gc.GamesService.End(&payloads.EndGamePayload{
		Id: id,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, game)
}