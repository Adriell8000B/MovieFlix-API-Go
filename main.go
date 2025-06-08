package main

import (
	"MovieFlix-API-Go/server"
	"MovieFlix-API-Go/utils"
	"net/http"
	"os"
)

func main() {
	utils.SetupEnviroment()

	server_conf := &http.Server{
		Addr: os.Getenv("PORT"),
		Handler: nil,
	}

	Server := server.NewServer(server_conf)

	Server.Init()
}