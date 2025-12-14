package main

import (
	"CardService/config"
	"CardService/internal/app"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	cfg := config.ConfigUP()

	s, err := app.NewApp(ctx, *cfg)
	if err != nil {
		log.Fatalf("Failed to create new App: %v", err)
	}

	go func() {
		if err := s.Run(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<- sig

	log.Println("Graceful shut Down")
	s.Stop()
}