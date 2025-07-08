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
	englishString := "task"
	cyrillicString := "главрыба"
	emojiString := "😁😎🤩"
	fmt.Printf("Входная строка: %s, перевернутая строка: %s\n", englishString, reverseString(englishString))
	fmt.Printf("Входная строка: %s, перевернутая строка: %s\n", cyrillicString, reverseString(cyrillicString))
	fmt.Printf("Входная строка: %s, перевернутая строка: %s\n", emojiString, reverseString(emojiString))
}
