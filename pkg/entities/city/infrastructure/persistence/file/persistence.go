package file

import (
	"bufio"
	"context"
	"github.com/kenesparta/paackgo/config"
	"github.com/kenesparta/paackgo/entities/city/domain"
	"log"
	"os"
)

type CityPersistenceFileStore struct {
	varConfig *config.VariableConfig
}

func NewCityPersistenceFileStore(v *config.VariableConfig) *CityPersistenceFileStore {
	return &CityPersistenceFileStore{
		varConfig: v,
	}
}

func (cp *CityPersistenceFileStore) Find(_ context.Context, cityId domain.CityId) (*domain.City, error) {
	fileOpen, err := os.Open(cp.varConfig.File.Name)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fileOpen)

	return nil, domain.ErrCityNotFound
}
