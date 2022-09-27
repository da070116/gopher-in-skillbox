package repository

import (
	"skillbox-test/pkg"
)

type CityCache struct {
	storage *pkg.StorageCache
}

func (r *CityCache) FilterCitiesByPopulation(min int, max int) ([]pkg.City, error) {
	return r.storage.FilterCitiesByPopulationRange(min, max), nil
}

func (r *CityCache) FilterCitiesByFoundation(min int, max int) ([]pkg.City, error) {
	return r.storage.FilterCitiesByFoundationRange(min, max), nil
}

func (r *CityCache) FilterCitiesByDistrict(district string) ([]pkg.City, error) {
	return r.storage.FilterCitiesByDistrict(district), nil
}

func NewCityCache(db *pkg.StorageCache) *CityCache {
	return &CityCache{storage: db}
}

func (r *CityCache) CreateCity(cityData string) (pkg.City, error) {
	return r.storage.SetCity(cityData)
}

func (r *CityCache) GetAllCities() ([]pkg.City, error) {
	return r.storage.GetAllCities(), nil
}

func (r *CityCache) FilterCitiesByRegion(region string) ([]pkg.City, error) {
	return r.storage.FilterCitiesByRegion(region), nil
}

func (r *CityCache) DeleteCity(deleteId int) error {
	return r.storage.DeleteCity(deleteId)
}

func (r *CityCache) UpdateCityPopulation(updateId int, data pkg.CityPopulation) error {
	value := *data.Population
	return r.storage.UpdateCityPopulation(updateId, value)
}
