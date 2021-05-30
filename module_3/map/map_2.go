package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   {"Деревня", "Село"},        // города с населением 10-99 тыс. человек
		100:  {"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: {"Миллионик"},              // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Село":          10,
		"cело":          100,
		"Миллионик":     1000,
		"Город":         100,
		"Большой город": 100,
	}

	// for key, cities := range groupCity {
	// 	for _, city := range cities {
	// 		if key == 100 {
	// 			continue
	// 		} else {
	// 			for k := range cityPopulation {
	// 				if k == city {
	// 					delete(cityPopulation, k)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	for key := range cityPopulation {
		var exist bool
		for _, city := range groupCity[100] {
			if city == key {
				exist = true
				break
			}
		}
		if !exist {
			delete(cityPopulation, key)
		}
	}

	fmt.Println(cityPopulation)
}
