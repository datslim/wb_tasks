package main

import "fmt"

// функция для переворота строки
func reverseString(inputString string) string {
	// для правильной обработки символов, размер которых не равен 1 байту и нам не известен, приведем входную строку
	// к типу []rune и будем менять порядок символов in-place
	stringRunes := []rune(inputString)

	// используем два указателя, один из которых указывает на начало, а второй на конец исходного среза
	// таким образом за один проход, мы можем переписать наш слайс на обратный порядок символов
	l := 0
	r := len(stringRunes) - 1

	for l <= r { // пока указатели не встретились
		stringRunes[l], stringRunes[r] = stringRunes[r], stringRunes[l] // меняем начало с концом
		l++                                                             // левый указатель двигаем вперед
		r--                                                             // правый указатель двигаем назад
	}
	return string(stringRunes) // приводим наш срез к строке
}

func main() {
	firstTest := "task"
	secondTest := "главрыба"
	fmt.Println(reverseString(firstTest))
	fmt.Println(reverseString(secondTest))
}
