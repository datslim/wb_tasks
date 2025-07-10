package main

import (
	"fmt"
	"sort"
	"wb_tech/misc"
)

func quickSort(inputArray []int) []int {
	if len(inputArray) <= 1 { // если слайс содержит 1 элемент или меньше, то он уже отсортирован и его можно вернуть
		return inputArray
	}

	// выбираем опорное значение (в зависимости от того, какой элемент средний по значению)
	// в случае, если все элементы равны, то выбираем первый элемент
	first, mid, last := inputArray[0], inputArray[len(inputArray)/2], inputArray[len(inputArray)-1]
	var pivotValue int

	if first <= mid && mid <= last {
		pivotValue = mid
	} else if first <= last && last <= mid {
		pivotValue = last
	} else {
		pivotValue = first
	}

	// создаем слайсы для хранения элементов меньше, больше и равных опорному значению
	// capacity выбираем исходя из соображений оптимизации памяти
	// отдавать под элементы равные опорному значения треть размера исходного массива,
	// может показаться избыточным, но это выгодно, если в массиве много дубликатов
	lessThanPivot := make([]int, 0, len(inputArray)/3) // предполагаем, что в треть элементов будет меньше опорного значения
	moreThanPivot := make([]int, 0, len(inputArray)/3) // предполагаем, что в треть элементов будет больше опорного значения
	equalToPivot := make([]int, 0, len(inputArray)/3)  // предполагаем, что треть элементов будет равна опорному

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

	// можно было бы использовать append, но тогда пришлось бы перевыделять память после добавления нового элемента
	sortedArray := make([]int, len(inputArray))                             // сразу создаем слайс для результата с длиной исходного слайса
	copy(sortedArray, lessThanPivot)                                        // копируем слайс с элементами меньше опорного
	copy(sortedArray[len(lessThanPivot):], equalToPivot)                    // копируем слайс с опорными элементами
	copy(sortedArray[len(lessThanPivot)+len(equalToPivot):], moreThanPivot) // копируем слайс с элементами больше опорного

	return sortedArray
}

func main() {
	// сортировка с использованием quickSort
	originalArr := make([]int, 10)
	misc.FillSetWithRandomNumbers(originalArr) // заполняем случайными числами

	arr := make([]int, len(originalArr)) // создаем слайс для сортировки с помощью sort.Ints
	copy(arr, originalArr)               // копируем исходный слайс в новый для сохранения исходного состояния

	quickSortResult := quickSort(arr)
	fmt.Printf("Массив до сортировки:	%v\nСортировка (quickSort): %v\n", originalArr, quickSortResult)

	// сортировка с использованием sort.Ints (для сравнения результата)
	sort.Ints(arr)
	fmt.Printf("Cортировка (sort.Ints): %v\n", arr)
}
