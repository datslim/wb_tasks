package main

import "fmt"

const (
	BIT = int64(1) // 1 в двоичной системе
)

// функция для установки или очистки бита в числе
// аргументы:
// number - указатель на число, в котором нужно установить или очистить бит
// position - позиция бита
// bit - значение бита (0 или 1)
func setBit(number *int64, position, bit int) {
	switch bit {
	case 0:
		// очищаем бит с помощью  инвертированной маски и логического AND
		mask := ^(BIT << position)
		*number = *number & mask
	case 1:
		// устанавливаем бит с помощью маски и логического OR
		mask := BIT << position
		*number = *number | mask
	}
}
func main() {
	var value int64 = 5 // 101 в двоичной системе
	fmt.Printf("Бинарное представление числа %d до установки первого бита в 0: %b\n", value, value)
	setBit(&value, 0, 0) // 100 в двоичной системе
	fmt.Printf("Бинарное представление числа 5 после установки первого бита в 0: %b (%d)\n", value, value)
	setBit(&value, 1, 1) // 110 в двоичной системе
	fmt.Printf("Бинарное представление числа 4 после установки второго бита в 1: %b (%d)\n", value, value)
}
