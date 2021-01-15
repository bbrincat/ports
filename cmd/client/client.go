package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main(){


	//TODO initialise repositories and services
	//portService := infrastructure.NewGRPCPortSenderService()
	//service.NewGRPCPortLoaderService()
	//portService

	//Mainly Boilerplate server code
	ctx, cancel := context.WithCancel(context.Background())


	//TODO initialise api with repositories and services as dependencies

	mux := http.NewServeMux()
	mux.HandleFunc("/port/{code}", func(w http.ResponseWriter, req *http.Request) {
		//TODO split this into separate api package to separate out of main.
	})

	httpServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	// Run server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}