package controllers

import (
	"main/games/presenters/http/dtos"
	"main/router"
	"main/router/utils"
	"net/http"
)

type GamesController struct{}

func (gc *GamesController) Create(writer http.ResponseWriter, request *http.Request) {
	var createGameDto dtos.CreateGameDto
	err := utils.ReadHttpRequestBody(writer, request, &createGameDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusCreated, createGameDto)
}

func (gc *GamesController) Start(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	utils.WriteHttpResponse(writer, http.StatusOK, id)
}

func (gc *GamesController) End(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	utils.WriteHttpResponse(writer, http.StatusOK, id)
}