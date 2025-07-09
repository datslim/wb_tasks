package main

import (
	"fmt"
	"math/rand"
)

// функция для нахождения пересечения двух неупорядочных множеств (слайсов)
// возвращаемое значение: множество пересечений
func setsIntersection(firstSet, secondSet []int) []int {
	seenLength := len(firstSet) + len(secondSet)             // длина карты в худшем случае равна сумме длины обоих множеств
	intersectionLength := min(len(firstSet), len(secondSet)) // длина выходного множества не может быть больше длины минимального входного множества
	// карта для отслеживания вхождений чисел, ключ - число, значение - видели ли такое число до этого
	seen := make(map[int]bool, seenLength)
	// слайс вхождений
	intersection := make([]int, 0, intersectionLength)

	// проходимся по первому множеству и проставляем флаги на найденые числа
	for _, value := range firstSet {
		seen[value] = true
	}

	// проходимся по второму множеству и если встречаем числа,
	// которые уже есть в словаре - добавляем в результатирующее множество
	for _, value := range secondSet {
		if seen[value] {
			intersection = append(intersection, value)
			seen[value] = false // убираем флаг, чтобы избежать дубликатов
		}

	}
	return intersection
}

// функция для случайного заполнения пяти значений для множества чисел
func randomFillSet(set []int) {
	for i := range 5 {
		randomValue := rand.Int() % 10
		set[i] = randomValue
	}
}

func main() {
	firstSet := make([]int, 5) // создаем множество размером 5
	secondSet := make([]int, 5)
	randomFillSet(firstSet) // заполняем случайными числами
	randomFillSet(secondSet)
	fmt.Printf("Первое множество: %v\n", firstSet)
	fmt.Printf("Второе множество: %v\n", secondSet)
	fmt.Printf("Пересечение множеств: %v\n", setsIntersection(firstSet, secondSet)) // получаем и выводим пересечение
}
