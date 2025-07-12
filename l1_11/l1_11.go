package main

import (
	"fmt"

	"github.com/datslim/wb-tech-misc/misc"
)

// функция для нахождения пересечения двух неупорядочных множеств (слайсов), содержащих любые данные
// возвращаемое значение: множество пересечений
func setsIntersection[T comparable](firstSet, secondSet []T) []T {
	seenLength := len(firstSet) + len(secondSet)             // длина словаря в худшем случае равна сумме длины обоих множеств
	intersectionLength := min(len(firstSet), len(secondSet)) // длина выходного множества не может быть больше длины минимального входного множества
	// карта для отслеживания вхождений чисел, ключ - элемент, значение - видели ли такой элемент до этого
	seen := make(map[T]bool, seenLength)
	// слайс вхождений
	intersection := make([]T, 0, intersectionLength)

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

func main() {
	firstSet := make([]int, 5) // создаем множество размером 5
	secondSet := make([]int, 5)
	misc.FillSetWithRandomNumbers(firstSet) // заполняем случайными числами
	misc.FillSetWithRandomNumbers(secondSet)
	numbersIntersection := setsIntersection(firstSet, secondSet)
	fmt.Printf("Первое множество чисел: %v\nВторое множество чисел: %v\nПересечение множеств чисел: %v\n\n", firstSet, secondSet, numbersIntersection)

	anySet1 := make([]any, 10) // создаем множество размером 10
	anySet2 := make([]any, 10)

	misc.FillSetWithRandomValues(anySet1) // заполняем случайными значениями
	misc.FillSetWithRandomValues(anySet2)
	anyIntersection := setsIntersection(anySet1, anySet2)
	fmt.Printf("Первое множество: %v\nВторое множество: %v\nПересечение множеств: %v\n", anySet1, anySet2, anyIntersection)
}
