package main

import (
	"fmt"
	"math/rand"
	"wb_tech/misc"
)

// функция для удаления  i-го элемента из слайса любого типа
func removeSliceElement[T any](inputArray *[]T, indexToRemove int) {
	// создаем новый срез, с новым нижележащим массивом (во избежание утечек памяти),
	// куда добавляем все элементы до i-го и после i-го
	*inputArray = append((*inputArray)[:indexToRemove], (*inputArray)[indexToRemove+1:]...)

}

func main() {
	sliceInt := make([]int, 10)             // создаем слайс 10 чисел
	misc.FillSetWithRandomNumbers(sliceInt) // заполняем случайными числами
	fmt.Printf("Исходный слайс: 		    %v\n", sliceInt)
	sliceIntRemoveIndex := rand.Intn(len(sliceInt) - 1) // выбираем случайный индекс для удаления
	removeSliceElement(&sliceInt, sliceIntRemoveIndex)  // удаляем элемент
	fmt.Printf("Слайс после удаления %d-го элемента: %v\n\n", sliceIntRemoveIndex, sliceInt)

	sliceString := make([]string, 10)
	misc.FillSetWithRandomStrings(sliceString) // заполняем случайными строками
	fmt.Printf("Исходный слайс: 		    %v\n", sliceString)
	sliceStringToRemoveIndex := rand.Intn(len(sliceString) - 1) // выбираем случайный индекс для удаления
	removeSliceElement(&sliceString, sliceStringToRemoveIndex)  // удаляем элемент
	fmt.Printf("Слайс после удаления %d-го элемента: %v\n", sliceStringToRemoveIndex, sliceString)

}
