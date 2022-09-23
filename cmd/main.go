package main

import (
	"fmt"
	"skillbox-test/pkg"
)

func main() {
	sc := pkg.NewStorageCache()
	filename := "cities.csv"
	sc.LoadFromFile(filename)

	sib := sc.FilterCitiesByDistrict("Сибирский")
	fmt.Println(sib)

	year_gt_1700 := sc.FilterCitiesByFoundationRange(1700, 2000)
	fmt.Println(year_gt_1700)

	sc.SetCity("1,Тестовск,Тестовский район,Задачный,0,2022")
	sc.LoadIntoFile("cities2.csv")
}
