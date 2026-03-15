package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaronlee232/redis-gui-tester/internal/app"
	"github.com/aaronlee232/redis-gui-tester/internal/database"
)

func main() {
	fmt.Println("Running redis-gui-tester")
	db := database.InitDB()

	defer db.Close()

	// Start Server
	addr := ":3000"
	router := app.NewRouter(db)

	log.Println("Starting server on port", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
