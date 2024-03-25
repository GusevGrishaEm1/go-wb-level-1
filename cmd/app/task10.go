package main

import "fmt"

func task10() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mp := make(map[int][]float64)
	for _, el := range array {
		val := (int(el) - (int(el) % 10))
		_, ok := mp[val]
		if ok {
			mp[val] = append(mp[val], el)
		} else {
			mp[val] = []float64{el}
		}
	}
	for key, val := range mp {
		fmt.Print(key, " --- ")
		for _, el := range val {
			fmt.Printf("%.1f ", el)
		}
		fmt.Print("\n")
	}
}