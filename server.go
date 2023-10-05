package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/arvind-prajapati/Go-Pass-Wallet/factory"
	"github.com/arvind-prajapati/Go-Pass-Wallet/internal"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load environment", "error", err)
		os.Exit(0)
	}

	port := factory.MustGetenv("PORT")

	router := mux.NewRouter()
	// Generate pass template
	router.HandleFunc("/generate/pass", internal.GeneratePass).Methods(http.MethodGet)

	// Register pass and its associated device in database.
	router.HandleFunc("/devices/{deviceLibraryIdentifier}/registrations/{passTypeIdentifier}/{serialNumber}", internal.RegisterPass).Methods(http.MethodPost)

	// Get the List of Updatable Passes
	router.HandleFunc("/devices/{deviceLibraryIdentifier}/registrations/{passTypeIdentifier}", internal.GetListUpdatablePasses).Methods(http.MethodGet)

	// Send an Updated Pass
	router.HandleFunc("/passes/{passTypeIdentifier}/{serialNumber}", internal.SendUpdatedPass).Methods(http.MethodGet)

	// Unregister a Pass
	router.HandleFunc("/devices/{deviceLibraryIdentifier}/registrations/{passTypeIdentifier}/{serialNumber}", internal.UnregisterPass).Methods(http.MethodDelete)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	slog.Info(fmt.Sprintf("server is listening on port:%s", port))
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
