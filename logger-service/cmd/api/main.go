package main

import (
	"log"
	"ms/logger-service/internal/data"
	"ms/logger-service/internal/server"
)

const (
	webPort  = ":80"
	rpcPort  = "5001"
	mongoUrl = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

func main() {
	_, err := data.ConnectToDB(mongoUrl)
	if err != nil {
		log.Panic("Couldn't connect to database...", err)
	}
	defer data.CloseConnection()

	srv := &server.Server{
		Models: data.New(),
	}

	srv.Run(webPort)
}
