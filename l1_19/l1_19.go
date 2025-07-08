package main

import "fmt"

// функция для переворота строки
func reverseString(inputString string) string {
	// для правильной обработки символов, размер которых не равен 1 байту и нам не известен, приведем входную строку
	// к типу []rune и будем менять порядок символов in-place
	inputRunes := []rune(inputString)

	// используем два указателя, один из которых указывает на начало, а второй на конец исходного среза
	// таким образом за один проход, мы можем переписать наш слайс на обратный порядок символов
	l := 0
	r := len(inputRunes) - 1

	for l <= r { // пока указатели не встретились
		inputRunes[l], inputRunes[r] = inputRunes[r], inputRunes[l] // меняем начало с концом
		l++                                                         // левый указатель двигаем вперед
		r--                                                         // правый указатель двигаем назад
	}
	return string(inputRunes) // приводим наш срез к строке
}

func main() {
	englishString := "task"
	cyrillicString := "главрыба"
	emojiString := "😁😎🤩"
	fmt.Printf("Входная строка: %s\nПеревернутая строка: %s\n", englishString, reverseString(englishString))
	fmt.Printf("Входная строка: %s\nПеревернутая строка: %s\n", cyrillicString, reverseString(cyrillicString))
	fmt.Printf("Входная строка: %s\nПеревернутая строка: %s\n", emojiString, reverseString(emojiString))
}
