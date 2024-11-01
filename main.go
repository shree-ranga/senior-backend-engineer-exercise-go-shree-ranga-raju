package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	appcontext "syndio/appContext"
	"syndio/db/migrations"
	"syndio/router"
	"syndio/services"
)

func main() {
	// Init app context
	appcontext.Init()

	// instantiate dependencies
	dependencies := services.InstantiateServerDependencies()

	// Init router
	r := router.InitRouter(*dependencies)

	// Get the server port from the environment or set a default
	port := os.Getenv("SYNDIO_PORT")
	if port == "" {
		port = "8000"
	}
	addr := fmt.Sprintf(":%s", port)

	// apply migrations if any
	err := migrations.Migrate()
	if err != nil {
		log.Fatal("Migration error: ", err.Error())
	}

	// create and start the http server
	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Printf("Server starting on port: %s", port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server startup failed: %v", err)
	}
}
