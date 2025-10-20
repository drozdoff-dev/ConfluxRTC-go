package memory

import (
	"errors"
	"room-service/internal/domain"
	"sync"
)

type InMemoryRoomRepository struct {
	mu    sync.RWMutex
	rooms map[string]*domain.Room
}

func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{
		rooms: make(map[string]*domain.Room),
	}
}

func (r *InMemoryRoomRepository) Create(room *domain.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.rooms[room.ID] = room
	return nil
}

func (r *InMemoryRoomRepository) GetByID(id string) (*domain.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	room, ok := r.rooms[id]
	if !ok {
		return nil, errors.New("room not found")
	}
	return room, nil
}
