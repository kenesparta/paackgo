package persistence

import (
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/entities/trip/domain"
	"github.com/kenesparta/paackgo/entities/trip/infrastructure/persistence/memory"
)

type TripInMemory map[domain.TripId]domain.Trip

// MemoryRepository Here you can put all the In-Memory repositories available in the project.
type MemoryRepository struct {
	Trip *memory.TripMemoryPersistence
}

// NewMemoryRepository Configures the in-memory repository
func NewMemoryRepository(v *config.VariableConfig) (*MemoryRepository, error) {
	tim := make(TripInMemory)
	return &MemoryRepository{
		Trip: memory.NewTripMemoryPersistence(v, tim),
	}, nil
}
