package main

import "fmt"

// функция для замены значений двух чисел без временной переменной используя сложение\вычитание
// функция принимает указатель на два числа, которые нужно поменять между собой
// first = 5, second = 9
// first = 5 + 9 = 14
// second = 14 - 9 = 5
// first = 14 - 5 = 9
// first = 9, second = 5
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
