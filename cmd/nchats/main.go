package main

import (
	"log"
	"os"
	"strconv"

	"github.com/blacksails/nchats"
	"github.com/nats-io/go-nats"
)

func main() {
	portStr := os.Getenv("PORT")
	port := 8080
	if portStr != "" {
		iport, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
		}
		port = iport
	}

	natsURL := os.Getenv("NATSURL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	err := nchats.Start(nchats.Options{
		Port:    port,
		NATSURL: natsURL,
	})
	if err != nil {
		log.Fatal(err)
	}
}
