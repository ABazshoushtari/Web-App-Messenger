package main

import (
	"github.com/ABazshoushtari/Web-App-Messenger/internal/config"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/server"
	"log"
)

func main() {
	if err := config.Load("config.yaml"); err != nil {
		log.Fatal("error loading config")
	}
	logger.Init()
	logger.Logger().Fatal(server.Start())
}
