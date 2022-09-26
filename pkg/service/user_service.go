package service

import (
	"skillbox-test/pkg"
	"skillbox-test/pkg/repository"
)

type CityService struct {
	repo repository.City
}

func (s *CityService) DeleteCity(deleteId int) error {
	return s.repo.DeleteCity(deleteId)
}

func NewCityService(repo repository.City) *CityService {
	return &CityService{repo: repo}
}

func (s *CityService) CreateCity(cityData string) (pkg.City, error) {
	return s.repo.CreateCity(cityData)
}

func (s *CityService) GetAllCities() ([]pkg.City, error) {
	return s.repo.GetAllCities()
}

func (s *CityService) UpdateCity(cityId int, population pkg.CityPopulation) error {
	return s.repo.UpdateCityPopulation(cityId, population)
}
