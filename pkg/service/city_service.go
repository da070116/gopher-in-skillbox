package service

import (
	"skillbox-test/pkg"
	"skillbox-test/pkg/repository"
)

type CityService struct {
	repo repository.City
}

func (s *CityService) FilterCitiesByDistrict(district string) ([]pkg.City, error) {
	return s.repo.FilterCitiesByDistrict(district)
}

func (s *CityService) FilterCitiesByPopulation(min int, max int) ([]pkg.City, error) {
	return s.repo.FilterCitiesByPopulation(min, max)
}

func (s *CityService) FilterCitiesByFoundation(min int, max int) ([]pkg.City, error) {
	return s.repo.FilterCitiesByFoundation(min, max)
}

func (s *CityService) FilterCitiesByRegion(region string) ([]pkg.City, error) {
	return s.repo.FilterCitiesByRegion(region)
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
