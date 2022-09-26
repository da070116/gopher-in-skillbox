package repository

import (
	"skillbox-test/pkg"
)

type City interface {
	CreateCity(cityData string) (pkg.City, error)
	GetAllCities() ([]pkg.City, error)
	DeleteCity(deleteId int) error
	UpdateCityPopulation(updateId int, data pkg.CityPopulation) error
}

type Repository struct {
	City
}

// NewRepository - constructor for database maintainer
func NewRepository(storage *pkg.StorageCache) *Repository {
	return &Repository{
		City: NewCityCache(storage),
	}
}
