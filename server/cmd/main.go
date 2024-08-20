package main

import (
	"log"

	"github.com/HenryGunadi/simple-chat-app/server/cmd/api"
	_ "github.com/lib/pq"
)

func main() {
	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatalf("error running server : %v", err)
	}
}