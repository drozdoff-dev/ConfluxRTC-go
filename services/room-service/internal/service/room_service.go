package service

import (
	"context"
	"fmt"
	"room-service/internal/domain"
	"room-service/internal/repository"
	"time"

	proto "github.com/drozdoff-dev/ConfluxRTC-libs/proto/room-service"

	"github.com/google/uuid"
)

type RoomService struct {
	proto.UnimplementedRoomServiceServer
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *proto.CreateRoomRequest) (*proto.RoomResponse, error) {
	room := &domain.Room{
		ID:        uuid.NewString(),
		Name:      req.Name,
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(room); err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}
	return &proto.RoomResponse{
		Id:        room.ID,
		Name:      room.Name,
		CreatedAt: room.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *RoomService) GetRoom(ctx context.Context, req *proto.GetRoomRequest) (*proto.RoomResponse, error) {
	room, err := s.repo.GetByID(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}
	return &proto.RoomResponse{
		Id:        room.ID,
		Name:      room.Name,
		CreatedAt: room.CreatedAt.Format(time.RFC3339),
	}, nil
}
