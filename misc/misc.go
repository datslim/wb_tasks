package misc

import "math/rand"

// функция для заполнения слайса int случайными числами
func FillSetWithRandomNumbers(set []int) {
	for i := range set {
		randomValue := rand.Int() % 10
		set[i] = randomValue
	}
}

// функция для заполнения слайса типа any случайными значениями
func FillSetWithRandomValues(set []any) {
	// возможные значения для записи
	possibleValues := []any{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		"apple", "banana", "orange", "l1", "webdev", "wb-tech", "cat", "dog",
		1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9,
		true, false,
		nil, "golang", "programming", 42, 3.1415,
	}

	for i := range set {
		randomIndex := rand.Intn(len(possibleValues))
		set[i] = possibleValues[randomIndex]
	}
}
