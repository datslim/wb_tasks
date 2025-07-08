package main

import (
	"fmt"
	"strings"
)

// функция для разворота порядка слов в предложении, не используя дополнительные срезы (in-place)
func reverseWordsOrder(inputString string) string {
	// создаем срез строк, состоящий из слов предложения, разделенных пробелами
	words := strings.Split(inputString, " ")

	// используем два указателя
	l := 0
	r := len(words) - 1

	for l <= r {
		words[l], words[r] = words[r], words[l] // меняем порядок слов (начало с концом, и так до центра)
		l++
		r--
	}

	// возвращаем строку, состоящую из полученного массив слов и разделяем их пробелом
	return strings.Join(words, " ")
}

func main() {
	englishString := "task is very hard"
	cyrillicString := "главная рыба главнее остальных"
	emojiString := "😁 😎 🤩"
	fmt.Printf("Входная строка: %s\nПеревернутая строка: %s\n", englishString, reverseWordsOrder(englishString))
	fmt.Printf("Входная строка: %s\nПеревернутая строка: %s\n", cyrillicString, reverseWordsOrder(cyrillicString))
	fmt.Printf("Входная строка: %s\nПеревернутая строка: %s\n", emojiString, reverseWordsOrder(emojiString))
}
