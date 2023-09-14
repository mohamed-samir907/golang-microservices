package main

import (
	"ms/broker-service/internal/server"
)

const webPort = ":80"

func main() {
	server.Run(webPort)
}
