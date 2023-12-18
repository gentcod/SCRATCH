package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gentcod/RSSAggregator/api"
	"github.com/gentcod/RSSAggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	//Configure Port
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbUrl := os.Getenv("DB_CONNECTION_STRING")
	if dbUrl == "" {
		log.Fatal("Database URL not found in the environment")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	log.Printf("Server starting on port %v", port)
	err = server.Start(port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", port)
}