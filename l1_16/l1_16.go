package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func quickSort(inputArray []int) []int {
	if len(inputArray) <= 1 { // если слайс содержит 1 элемент или меньше, то он уже отсортирован и его можно вернуть
		return inputArray
	}

	pivotValue := inputArray[len(inputArray)/2] // получаем опорное значение (средний элемент)

	// создаем слайсы для хранения элементов меньше, больше и равных опорному значению
	// capacity выбираем исходя из соображений оптимизации памяти
	lessThanPivot := make([]int, 0, len(inputArray)/2) // предполагаем, что в среднем половина элементов будет меньше опорного значения
	moreThanPivot := make([]int, 0, len(inputArray)/2) // предполагаем, что в среднем половина элементов будет больше опорного значения
	equalToPivot := make([]int, 0, 1)                  // предполагаем, что опорный элемент встретится только один раз

	// разбиваем исходный слайс на три части: меньше, больше и равных опорному значению и добавляем их в соответствующие слайсы
	for _, element := range inputArray {
		if element < pivotValue {
			lessThanPivot = append(lessThanPivot, element)
		} else if element > pivotValue {
			moreThanPivot = append(moreThanPivot, element)
		} else {
			equalToPivot = append(equalToPivot, element)
		}
	}

	// рекурсивно сортируем полученные слайсы
	lessThanPivot = quickSort(lessThanPivot)
	moreThanPivot = quickSort(moreThanPivot)

	// объединяем полученные слайсы в один
	return append(append(lessThanPivot, equalToPivot...), moreThanPivot...)
}

// функция для заполнения слайса случайными числами
func fillSetWithRandomNumbers(set []int) {
	for i := range set {
		randomValue := rand.Int() % 10
		set[i] = randomValue
	}
}

func main() {
	// сортировка с использованием quickSort
	originalArr := make([]int, 10)
	fillSetWithRandomNumbers(originalArr) // заполняем случайными числами

	arr := make([]int, len(originalArr)) // создаем слайс для сортировки с помощью sort.Ints
	copy(arr, originalArr)               // копируем исходный слайс в новый для сохранения исходного состояния

	quickSortResult := quickSort(arr)
	fmt.Printf("Массив до сортировки:	%v\nСортировка (quickSort): %v\n", originalArr, quickSortResult)

	// сортировка с использованием sort.Ints (для сравнения результата)
	sort.Ints(arr)
	fmt.Printf("Cортировка (sort.Ints): %v\n", arr)
}
