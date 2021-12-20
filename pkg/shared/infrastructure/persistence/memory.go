package persistence

import (
	"github.com/kenesparta/paackgo/config"
)

// MemoryRepository Here you can put all the In-Memory repositories available in the project.
type MemoryRepository struct {
}

// NewMemoryRepository Configures the in-memory repository
func NewMemoryRepository(v *config.VariableConfig) (*MemoryRepository, error) {
	return &MemoryRepository{}, nil
}
