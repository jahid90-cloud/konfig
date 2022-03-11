package main

import (
	"log"

	"github.com/jahid90-cloud/konfig/service/server/http"
	"github.com/jahid90-cloud/konfig/service/utils/env"
)

const ADDR string = ":8080"

func main() {
	env := env.NewEnv()
	httpServer := http.NewServer(env)

	var address string
	address = env.Get("ADDR")
	if len(address) == 0 {
		address = ADDR
	}

	if err := httpServer.Start(address); err != nil {
		log.Fatal(err)
	}
}
