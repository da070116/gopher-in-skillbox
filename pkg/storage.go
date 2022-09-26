package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
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
		_, err := sc.SetCity(scanner.Text())
		if err != nil {
			return err
		}
	}

	return nil
}

func (sc *StorageCache) LoadIntoFile(filename string) {

	file, err := os.Create(filename)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer CloseFile(file)

	for id, city := range sc.cities {
		cityString := fmt.Sprintf("%d,%s\n", id, CityWithIdToString(city))
		_, err := file.WriteString(cityString)
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}

func (sc *StorageCache) SetCity(rawData string) (City, error) {
	data := strings.Trim(rawData, " ")
	params := strings.Split(data, ",")

	uid, err := strconv.Atoi(params[0])
	if err != nil {
		return City{}, err
	}
	name := params[1]
	region := params[2]
	district := params[3]
	population, err := strconv.Atoi(params[4])
	if err != nil {
		return City{}, err
	}
	foundation, err := strconv.Atoi(params[5])
	if err != nil {
		return City{}, err
	}

	sc.Lock()
	defer sc.Unlock()

	newCity := City{
		Id:         uid,
		Name:       name,
		Region:     region,
		District:   district,
		Population: population,
		Foundation: foundation,
	}

	sc.cities[uid] = newCity

	return newCity, nil
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

func (sc *StorageCache) GetAllCities() []City {
	result := make([]City, 0)
	for _, city := range sc.cities {

		result = append(result, city)
	}
	return result
}
