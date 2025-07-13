package main

import (
	"log"

	"github.com/kyomel/go-scalable-educative/server"
)

func main() {

	var port = "8080"

	server := server.NewServer(port)
	err := server.PreStart()
	if err != nil {
		log.Fatalf("Error in prestart - %v\n", err)
	}

	done, err := server.Start()
	if err != nil {
		server.Shutdown()
		log.Fatalf("Error starting server - %v\n", err)
	}

	<-done
}
