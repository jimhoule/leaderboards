package main

import (
	"fmt"
	"log"
	"main/games"
	"main/players"
	"main/router"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Envs
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading envs", err)
	}

	// Router
	mainRouter := router.Get()

	// Modules
	games.Init(mainRouter)
	players.Init(mainRouter)

	// Http server
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("HTTP_URL"), os.Getenv("HTTP_PORT")),
		Handler: mainRouter,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}