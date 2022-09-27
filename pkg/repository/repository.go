package repository

import (
	"skillbox-test/pkg"
)

type City interface {
	CreateCity(cityData string) (pkg.City, error)
	GetAllCities() ([]pkg.City, error)
	DeleteCity(deleteId int) error
	UpdateCityPopulation(updateId int, data pkg.CityPopulation) error
	FilterCitiesByRegion(region string) ([]pkg.City, error)
	FilterCitiesByDistrict(district string) ([]pkg.City, error)
	FilterCitiesByPopulation(min int, max int) ([]pkg.City, error)
	FilterCitiesByFoundation(min int, max int) ([]pkg.City, error)
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
