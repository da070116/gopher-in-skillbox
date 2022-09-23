package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type StorageCache struct {
	sync.RWMutex
	cities map[int]City
}

func NewStorageCache() *StorageCache {
	cities := make(map[int]City)

	storageCache := StorageCache{
		cities: cities,
	}

	return &storageCache
}

func (sc *StorageCache) LoadFromFile(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sc.SetCity(scanner.Text())
	}

	return nil
}

func (sc *StorageCache) LoadIntoFile(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer CloseFile(file)

	for id, city := range sc.cities {
		cityString := fmt.Sprintf("%d,%s\n", id, sc.PutCity(city))
		file.WriteString(cityString)
	}

	return nil
}

func (sc *StorageCache) PutCity(city City) string {
	stringvals := []string{city.Name, city.Region, city.District, strconv.Itoa(city.Population), strconv.Itoa(city.Foundation)}
	return strings.Join(stringvals, ",")
}

func (sc *StorageCache) SetCity(rawData string) error {
	data := strings.Trim(rawData, " ")
	params := strings.Split(data, ",")

	uid, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	name := params[1]
	region := params[2]
	district := params[3]
	population, err := strconv.Atoi(params[4])
	if err != nil {
		return err
	}
	foundation, err := strconv.Atoi(params[5])
	if err != nil {
		return err
	}

	sc.Lock()
	defer sc.Unlock()
	sc.cities[uid] = City{
		Name:       name,
		Region:     region,
		District:   district,
		Population: population,
		Foundation: foundation,
	}

	return nil
}

func (sc *StorageCache) GetCity(uid int) (City, error) {
	sc.RLock()
	defer sc.RUnlock()

	city, ok := sc.cities[uid]
	if !ok {
		return City{}, errors.New("no such data")
	}
	return city, nil
}

func (sc *StorageCache) DeleteCity(uid int) error {
	sc.Lock()
	defer sc.Unlock()

	if _, found := sc.cities[uid]; !found {
		return errors.New("no such data")
	}
	delete(sc.cities, uid)
	return nil
}

func (sc *StorageCache) UpdateCityPopulation(uid int, population int) error {
	city, err := sc.GetCity(uid)
	if err != nil {
		return err
	}
	city.Population = population

	sc.cities[uid] = city
	return nil
}

func (sc *StorageCache) FilterCitiesByRegion(region string) []City {
	result := make([]City, 0)
	for _, city := range sc.cities {
		if city.Region == region {
			result = append(result, city)
		}
	}
	return result
}

func (sc *StorageCache) FilterCitiesByDistrict(district string) []City {
	result := make([]City, 0)
	for _, city := range sc.cities {
		if city.District == district {
			result = append(result, city)
		}
	}
	return result
}

func (sc *StorageCache) FilterCitiesByPopulationRange(min int, max int) []City {
	result := make([]City, 0)
	for _, city := range sc.cities {
		if city.Population >= min && city.Population <= max {
			result = append(result, city)
		}
	}
	return result
}

func (sc *StorageCache) FilterCitiesByFoundationRange(min int, max int) []City {
	result := make([]City, 0)
	for _, city := range sc.cities {
		if city.Foundation >= min && city.Foundation <= max {
			result = append(result, city)
		}
	}
	return result
}
