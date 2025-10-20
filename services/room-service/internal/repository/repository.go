package repository

import "room-service/internal/domain"

type RoomRepository interface {
	Create(room *domain.Room) error
	GetByID(id string) (*domain.Room, error)
}
