package persistence

import (
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/entities/city/infrastructure/persistence/file"
)

// FileRepository Here you can put all the File repositories available in the project.
type FileRepository struct {
	City *file.CityPersistenceFileStore
}

// NewFileRepository Configures the file repository
func NewFileRepository(v *config.VariableConfig) (*FileRepository, error) {
	return &FileRepository{
		City: file.NewCityPersistenceFileStore(v),
	}, nil
}
