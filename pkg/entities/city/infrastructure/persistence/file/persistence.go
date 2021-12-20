package file

import (
	"bufio"
	"context"
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/entities/city/domain"
	"os"
)

type CityFileStorePersistence struct {
	varConfig *config.VariableConfig
}

func NewCityFileStorePersistence(v *config.VariableConfig) *CityFileStorePersistence {
	return &CityFileStorePersistence{
		varConfig: v,
	}
}

func (cp *CityFileStorePersistence) Find(_ context.Context, cityId domain.CityId) (*domain.City, error) {
	fileOpen, err := os.Open(cp.varConfig.File.Name)
	if err != nil {
		return nil, err
	}

	countId := domain.CityId(1)
	scanner := bufio.NewScanner(fileOpen)
	for scanner.Scan() {
		if cityId == countId {
			return &domain.City{Id: cityId, Name: scanner.Text()}, nil
		}
		countId++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(fileOpen)

	return nil, domain.ErrCityNotFound
}
