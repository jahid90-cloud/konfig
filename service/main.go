package main

import (
	"log"

	"github.com/jahid90-cloud/konfig/service/server/http"
)

const ADDR string = ":8080"

func main() {
	httpServer := http.NewServer()

	if err := httpServer.Start(ADDR); err != nil {
		log.Fatal(err)
	}
}
