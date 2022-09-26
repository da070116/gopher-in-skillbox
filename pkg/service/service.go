package service

import (
	"skillbox-test/pkg"
	"skillbox-test/pkg/repository"
)

type City interface {
	CreateCity(cityData string) (pkg.City, error)
	GetAllCities() ([]pkg.City, error)
	DeleteCity(deleteId int) error
	UpdateCity(updateId int, data pkg.CityPopulation) error
}

type Service struct {
	City
}

// NewService constructor
func NewService(repo *repository.Repository) *Service {
	return &Service{
		City: NewCityService(repo),
	}
}
