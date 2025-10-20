package main

import (
	"log"
	"room-service/internal/app"
	"room-service/internal/repository/memory"
)

func main() {
	repo := memory.NewInMemoryRoomRepository()
	server := app.New(repo, 50051)

	if err := server.Run(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
