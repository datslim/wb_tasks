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
	// карта для отслеживания вхождений чисел, ключ - число, значение - видели ли такой элемент до этого
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

// функция для случайного заполнения слайса, пятью числами
func randomFillSet(set []any) {
	for i := range 5 {
		randomValue := rand.Int() % 10
		set[i] = randomValue
	}
}

func main() {
	firstSet := make([]any, 5) // создаем множество размером 5
	secondSet := make([]any, 5)
	randomFillSet(firstSet) // заполняем случайными числами
	randomFillSet(secondSet)
	fmt.Printf("Первое множество чисел: %v\n", firstSet)
	fmt.Printf("Второе множество чисел : %v\n", secondSet)
	fmt.Printf("Пересечение множеств чисел: %v\n", setsIntersection(firstSet, secondSet)) // получаем и выводим пересечение

	anySet1 := []any{1, "hello", 3.14, true, "world", "golang"}
	anySet2 := []any{2, "hello", 3.14, true, "test", "golang"}

	fmt.Printf("Первое множество: %v\n", anySet1)
	fmt.Printf("Второе множество: %v\n", anySet2)
	fmt.Printf("Пересечение множеств: %v\n", setsIntersection(anySet1, anySet2))
}
