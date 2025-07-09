package main

import (
	"fmt"
)

// функция для группировки массива температур на подгруппы, кратные 10
// -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5 ->
// -> -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}
func groupTemperatures(temperatures []float64) map[int][]float64 {
	groupedTemperatures := make(map[int][]float64, len(temperatures))

	for _, value := range temperatures {
		// -25.4 | -25.4 / 10 = -2.5 | int(-2.5) = -2 | -2 * 10 = -20
		group := int(value/10) * 10 // получаем группу температур
		groupedTemperatures[group] = append(groupedTemperatures[group], value)
	}

	return groupedTemperatures
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(groupTemperatures(temperatures))
}
