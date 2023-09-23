package main

import (
	"ms/mail-service/internal/server"
)

const (
	webPort = ":80"
)

func main() {
	srv := server.New()

	srv.Run(webPort)
}
