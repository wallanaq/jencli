package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var httpServer *http.Server

func Start(addr string, handler http.Handler) error {
	httpServer = &http.Server{Addr: addr, Handler: handler}

	log.Printf("Starting server at %s\n", addr)

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("ListenAndServe(): %v", err)
	}

	return nil

}

func Shutdown() error {
	if httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			return fmt.Errorf("error when shutting down the server: %v", err)
		}

		log.Println("Server gracefully stopped")
	}

	return nil
}
