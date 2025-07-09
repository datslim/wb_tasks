package main

import "fmt"

// функция для нахождения пересечения двух неупорядочных множеств (слайсов)
// возвращаемое значение: множество пересечений
func setsIntersection(firstSet, secondSet []int) []int {
	// карта для отслеживания вхождений чисел, ключ - число, значение - видели ли такое число до этого
	seen := make(map[int]bool, 0)
	// слайс вхождений
	intersection := make([]int, 0)

	// проходимся по первому множеству и проставляем флаги на найденые числа
	for _, value := range firstSet {
		seen[value] = true
	}

	// проходимся по второму множеству и если встречаем числа,
	// которые уже есть в словаре - добавляем в результатирующее множество
	for _, value := range secondSet {
		if seen[value] == true {
			intersection = append(intersection, value)
		}

	}
	return intersection
}

func main() {
	firstSet := []int{1, 2, 3, 1}
	secondSet := []int{2, 3, 4}
	fmt.Println(setsIntersection(firstSet, secondSet))
}
