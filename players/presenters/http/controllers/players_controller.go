package controllers

import (
	"main/players/application/payloads"
	"main/players/application/services"
	"main/players/presenters/http/dtos"
	"main/router"
	"main/router/utils"
	"net/http"
)

type PlayersController struct{
	PlayersService *services.PlayersService
}

func (pc *PlayersController) GetAll(writer http.ResponseWriter, request *http.Request) {
	players, err := pc.PlayersService.GetAll()
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, players)
}

func (pc *PlayersController) GetById(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	player, err := pc.PlayersService.GetById(id)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, player)
}

func (pc *PlayersController) Create(writer http.ResponseWriter, request *http.Request) {
	var createPlayerDto dtos.CreatePlayerDto
	err := utils.ReadHttpRequestBody(writer, request, &createPlayerDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	player, err := pc.PlayersService.Create(&payloads.CreatePlayerPayload{
		Name: createPlayerDto.Name,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusCreated, player)
}