package main

import (
	"fmt"
)

// функция для бинарного поиска числа с помощью двух указателей
// возвращаемое значение: индекс первого найденного элемента подходящего под условие, либо -1, если элемент не найден
func binarySearch(inputArray []int, target int) int {
	// создаем два указателя, левый - на начало слайса, правый - на его конец
	l, r := 0, len(inputArray)-1
	// идем до тех пор пока указатели не встретятся
	for l <= r {
		mid := (l + r) / 2             // находим середину текущего отрезка (между левым и правым указателем)
		if inputArray[mid] == target { // если середина - это искомое число, возвращаем его
			return mid
		}
		if inputArray[mid] > target { // если число посередине больше нужного нам
			r = mid - 1 // ставим правый указатель до середины
		} else { // если число посередине меньше нужного нам
			l = mid + 1 // ставим левый указатель после середины
		}
	}
	return -1
}

// функция бинарного поиска первого вхождения числа в массив
// возвращаемое значение: индекс первого элемента подходящего под условие, либо -1, если элемент не найден
func binarySearchFirst(inputArray []int, target int) int {
	l, r := 0, len(inputArray)-1
	result := -1

	for l <= r {
		mid := (l + r) / 2
		if inputArray[mid] == target {
			result = mid
			r = mid - 1 // продолжаем искать левее
		} else if inputArray[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return result
}

// функция бинарного поиска последнего вхождения числа в массив
// возвращаемое значение: индекс последнего элемента подходящего под условие, либо -1, если элемент не найден
func binarySearchLast(inputArray []int, target int) int {
	l, r := 0, len(inputArray)-1
	result := -1

	for l <= r {
		mid := (l + r) / 2
		if inputArray[mid] == target {
			result = mid
			l = mid + 1 // продолжаем искать правее
		} else if inputArray[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return result
}

func main() {
	array := []int{1, 2, 2, 2, 3, 3, 4, 5, 5, 5, 6, 7, 8, 8, 9} // создаем массив с дубликатами

	target := 2                                           // искомое число
	binarySearchAny := binarySearch(array, target)        // любой индекс вхождения числа
	binarySearchFirst := binarySearchFirst(array, target) // индекс первого вхождения числа
	binarySearchLast := binarySearchLast(array, target)   // индекс последнего вхождения числа

	fmt.Printf("Массив: %v\nЛюбое вхождение числа %d: индекс %d\nПервое вхождение числа %d: индекс %d\nПоследнее вхождение числа %d: индекс %d\n", array, target, binarySearchAny, target, binarySearchFirst, target, binarySearchLast)
}
