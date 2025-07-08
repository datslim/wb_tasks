package main

import "fmt"

// функция для замены значений двух чисел без временной переменной используя сложение\вычитание
// функция принимает указатель на два числа, которые нужно поменять между собой
// a = 5, b = 9
// a = 5 + 9 = 14
// b = 14 - 9 = 5
// a = 14 - 5 = 9
// a = 9, b = 5
func changeTwoInts(first, second *int) {
	*first += *second
	*second = *first - *second
	*first -= *second
}

func main() {
	first, second := 5, 9
	fmt.Printf("Перед заменой\nПервое значение: %d, второе значение: %d\n", first, second)
	changeTwoInts(&first, &second)
	fmt.Printf("После замены\nПервое значение: %d, второе значение: %d\n", first, second)
}
