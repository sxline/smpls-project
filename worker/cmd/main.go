package main

import (
	"github.com/sxline/smpls-project/worker/internal/constants"
	"github.com/sxline/smpls-project/worker/internal/service"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	// Exit the application with status code 0 (success)
	os.Exit(0)
}

// run initializes and runs the worker service.
func run() error {
	// Create a new instance of the worker service
	workerService := service.NewWorkerService()

	// Run the worker service, which involves starting goroutines for processing data
	workerService.Run(constants.JSONFilePath)

	// Return nil to indicate success
	return nil
}
