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
	var endGameDto dtos.EndGameDto
	err := utils.ReadHttpRequestBody(writer, request, &endGameDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	id := router.GetUrlParam(request, "id")
	game, err := gc.GamesService.End(&payloads.EndGamePayload{
		Id: id,
		WinnerId: endGameDto.WinnerId,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, game)
}

func (gc *GamesController) Join(writer http.ResponseWriter, request *http.Request) {
	var joinGameDto dtos.JoinGameDto
	err := utils.ReadHttpRequestBody(writer, request, &joinGameDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	id := router.GetUrlParam(request, "id")
	game, err := gc.GamesService.Join(&payloads.JoinGamePayload{
		Id: id,
		PlayerId: joinGameDto.PlayerId,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, game)
}