package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	papagoID := os.Getenv("PAPAGO_ID")
	papagoKey := os.Getenv("PAPAGO_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := new(http.ServeMux)

	photoListRoute(mux, apiKey)
	photoSearchRoute(mux, apiKey, papagoID, papagoKey)

	log.Println("server running on port 8080")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
