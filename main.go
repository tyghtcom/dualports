package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tyghtcom/dualports/handler"
)

func main() {
	publicSrv := &http.Server{
		Addr:    ":8080",
		Handler: &handler.PublicHandler{}, // Only allows GET /
	}

	internalMux := http.NewServeMux()
	internalMux.HandleFunc("POST /create", handler.Create)
	internalMux.HandleFunc("GET /read/", handler.Read)
	internalMux.HandleFunc("PUT /update/", handler.Update)
	internalMux.HandleFunc("DELETE /delete/", handler.Delete)

	internalSrv := &http.Server{
		Addr:    ":8081",
		Handler: internalMux,
	}

	// Handle OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start servers
	go func() {
		log.Println("Starting internal server on :8081")
		if err := internalSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Internal server error: %v", err)
		}
	}()

	go func() {
		log.Println("Starting public server on :8080")
		if err := publicSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Public server error: %v", err)
		}
	}()

	// Wait for termination signal
	<-quit
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown internal server
	if err := internalSrv.Shutdown(ctx); err != nil {
		log.Printf("Internal server shutdown error: %v", err)
	} else {
		log.Println("Internal server shut down gracefully")
	}

	// Shutdown public server
	if err := publicSrv.Shutdown(ctx); err != nil {
		log.Printf("Public server shutdown error: %v", err)
	} else {
		log.Println("Public server shut down gracefully")
	}
}
