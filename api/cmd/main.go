package main

import (
	"api/internal/config"
	"api/internal/grpc"
	"api/internal/server"
	serv "api/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// Run the application and handle any potential errors.
	if err := run(); err != nil {
		log.Fatal(err)
	}

	// Exit the application with a success status code.
	os.Exit(0)
}

// run initializes the configuration, sets up gRPC client, creates service instances,
// configures HTTP routes, and starts the HTTP server.
func run() error {
	// Get application configuration.
	cfg := config.GetConfig()
	grpcCfg := config.GetGrpcConfig()

	// Initialize the gRPC read client.
	grpcReadClient, err := grpc.NewGrpcReadClient(grpcCfg)
	if err != nil {
		return err
	}

	// Create the data service with the gRPC client.
	readService := serv.NewDataService(grpcReadClient)

	// Create an HTTP server with the data service.
	httpServer := server.NewHttpServer(readService)

	// Create a Gorilla mux router.
	router := mux.NewRouter()
	// Define routes and their corresponding HTTP handlers.
	router.HandleFunc("/api/v1/data", httpServer.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/statistic", httpServer.GetStatistic).Methods("GET")

	// Create an HTTP server instance.
	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	// Log the start of the API server.
	log.Printf("API started on %s", cfg.HTTPAddr)

	// Start the HTTP server and handle any potential errors.
	if err = srv.ListenAndServe(); err != nil {
		return err
	}

	// Return nil to indicate a successful execution.
	return nil
}
