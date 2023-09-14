package main

import (
	"log"
	"ms/auth-service/internal/data"
	"ms/auth-service/internal/server"
)

const webPort = ":80"

func main() {
	conn := data.ConnectToDB()
	if conn == nil {
		log.Panic("Couldn't connect to database...")
	}
	defer conn.Close()

	srv := &server.Server{
		DB:     conn,
		Models: data.New(conn),
	}

	srv.Run(webPort)
}
