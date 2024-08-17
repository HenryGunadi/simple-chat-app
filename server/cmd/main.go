package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/HenryGunadi/simple-chat-app/server/cmd/api"
	"github.com/HenryGunadi/simple-chat-app/server/config"
	"github.com/HenryGunadi/simple-chat-app/server/db"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.NewPostgreSQLStorage(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Envs.DBAddress, config.Envs.DBPort, config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName))
	if err != nil {
		log.Fatalf("db connection error : %v", err)
	}

	defer dbConn.Close()

	initStorage(dbConn)

	server := api.NewAPIServer(":8080", dbConn)
	if err := server.Run(); err != nil {
		log.Fatalf("error running server : %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to database: %s", config.Envs.DBName)
}