package misc

import "math/rand"

// функция для заполнения слайса случайными числами
func fillSetWithRandomNumbers(set []int) {
	for i := range set {
		randomValue := rand.Int() % 10
		set[i] = randomValue
	}
}
