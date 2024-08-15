package main

import (
	"fmt"
	"log"
	"main/games"
	"main/router"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading envs", err)
	}

	mainRouter := router.Get()

	games.Init(mainRouter)

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("HTTP_URL"), os.Getenv("HTTP_PORT")),
		Handler: mainRouter,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}