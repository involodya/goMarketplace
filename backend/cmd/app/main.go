package main

import (
	handler "fullstack/backend/internal/handler/http"
	"log"
	"os"
)

func main() {
	dbUri, ok := os.LookupEnv("DB_URI")
	if !ok {
		log.Println("cannot get DB_URI from ENV")
		dbUri = "main.db"
	}
	signingKey, ok := os.LookupEnv("AUTH_SIGNING_KEY")
	if !ok {
		log.Println("cannot get AUTH_SIGNING_KEY from ENV")
		signingKey = "my-signing-key"
	}

	handler.StartServer(":8000", dbUri, signingKey)
}
