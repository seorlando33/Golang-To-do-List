package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"toDoList/api"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	serverChan := make(chan os.Signal, 1)
	signal.Notify(serverChan, os.Interrupt, syscall.SIGTERM)

	server := api.New(os.Getenv("Port"))

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Server Started")

	<-serverChan

	server.Shutdown(ctx)
	log.Println("Server Stopped")
}
