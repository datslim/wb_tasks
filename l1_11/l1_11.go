package main

import (
	"fmt"
	"math/rand"
)

// функция для нахождения пересечения двух неупорядочных множеств (слайсов), содержащих любые данные
// возвращаемое значение: множество пересечений
func setsIntersection(firstSet, secondSet []any) []any {
	seenLength := len(firstSet) + len(secondSet)             // длина словаря в худшем случае равна сумме длины обоих множеств
	intersectionLength := min(len(firstSet), len(secondSet)) // длина выходного множества не может быть больше длины минимального входного множества
	// карта для отслеживания вхождений чисел, ключ - элемент, значение - видели ли такой элемент до этого
	seen := make(map[any]bool, seenLength)
	// слайс вхождений
	intersection := make([]any, 0, intersectionLength)

	// проходимся по первому множеству и проставляем флаги на найденые элементы
	for _, value := range firstSet {
		seen[value] = true
	}

	// проходимся по второму множеству и если встречаем элементы,
	// которые уже есть в словаре - добавляем в результатирующее множество
	for _, value := range secondSet {
		if seen[value] {
			intersection = append(intersection, value)
			seen[value] = false // убираем флаг, чтобы избежать дубликатов
		}

	}
	return intersection
}

// функция для заполнения слайса случайного числами
func fillSetWithRandomNumbers(set []any) {
	for i := range set {
		randomValue := rand.Int() % 10
		set[i] = randomValue
	}
}

// функция для заполнения слайса случайными значениями
func fillSetWithRandomValues(set []any) {
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

func main() {
	firstSet := make([]any, 5) // создаем множество размером 5
	secondSet := make([]any, 5)
	fillSetWithRandomNumbers(firstSet) // заполняем случайными числами
	fillSetWithRandomNumbers(secondSet)
	fmt.Printf("Первое множество чисел: %v\nВторое множество чисел: %v\nПересечение множеств чисел: %v\n\n", firstSet, secondSet, setsIntersection(firstSet, secondSet))

	anySet1 := make([]any, 10) // создаем множество размером 10
	anySet2 := make([]any, 10)

	fillSetWithRandomValues(anySet1) // заполняем случайными значениями
	fillSetWithRandomValues(anySet2)
	fmt.Printf("Первое множество: %v\nВторое множество: %v\nПересечение множеств: %v\n", anySet1, anySet2, setsIntersection(anySet1, anySet2))
}
