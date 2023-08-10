package main

import (
	"context"
	"fmt"
	"github/meshachdamilare/trimly/api/router"
	pgdb "github/meshachdamilare/trimly/repository/storage/postgres"
	"github/meshachdamilare/trimly/settings/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.Setup()
	pgdb.ConnectToDB()
}

func main() {
	//Load settings

	e := router.Setup()

	// The HTTP Server
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.Config.ServerPort),
		Handler: e,
	}

	// Server run context
	serverCtx, serverCancel := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownCancel := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		shutdownCancel()
		serverCancel()
	}()

	// Run the server
	fmt.Println("Server running on port: ", config.Config.ServerPort)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
