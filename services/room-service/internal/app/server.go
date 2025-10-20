package app

import (
	"fmt"
	"log"
	"net"

	"room-service/internal/repository"
	"room-service/internal/service"

	proto "github.com/drozdoff-dev/ConfluxRTC-libs/proto/room-service"
	"google.golang.org/grpc"
)

type App struct {
	repo repository.RoomRepository
	port int
}

func New(repo repository.RoomRepository, port int) *App {
	return &App{repo: repo, port: port}
}

func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	roomSrv := service.NewRoomService(a.repo)
	proto.RegisterRoomServiceServer(grpcServer, roomSrv)

	log.Printf("Room Service running on port %d", a.port)
	return grpcServer.Serve(lis)
}
