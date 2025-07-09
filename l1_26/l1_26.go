package main

import (
	"fmt"
)

// функция для проверки на то, что символы в строке уникальны (регистр не важен)
// пример: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false (повторяются a\a)
func uniqueSymbols(inputString string) bool {
	seen := make(map[rune]bool)               // создаем словарь, где ключ - это символ, а значение - встречался ли такой символ
	lowerRegistr := myOwnToLower(inputString) // переводим все символы в один регистр (нижний)

	for _, symbol := range lowerRegistr { // итерируемся по символам строки в нижнем регистре
		if seen[symbol] { // если символ уже true, значит символ не уникальный и можно вернуть false
			return false
		}
		seen[symbol] = true // отмечаем символ как встреченный
	}

	return true // возвращаем true, если все символы уникальны
}

// функция для приведения символов кириллицы и латиницы в нижний регистр
// можно было бы использовать функцию ToLower из библиотеки strings,
// но для общего развития я написал свою урезанную (из-за поддержки всего 2 алфавитов)
// версию этой функции
func myOwnToLower(inputString string) string {
	inputRunes := []rune(inputString)
	latinStep := 'A' - 'a'   // латинские А и а
	cyrllicStep := 'А' - 'а' // кириллические А и а
	for _, rune := range inputRunes {
		switch {
		case rune >= 'A' && rune <= 'Z':
			rune += latinStep
		case rune >= 'А' && rune <= 'Я':
			rune += cyrllicStep
		}
	}
	return string(inputRunes)
}

// небольшие тесты
func main() {
	var (
		test1 = "abcd"
		test2 = "abCdefAaf"
		test3 = "aabcd"
		test4 = "абуБАКА"
		test5 = "aBCD"
		test6 = "абвгд"
	)
	fmt.Printf("Результат строки abcd: %v\n", uniqueSymbols(test1))
	fmt.Printf("Результат строки abCdefAaf: %v\n", uniqueSymbols(test2))
	fmt.Printf("Результат строки aabcd: %v\n", uniqueSymbols(test3))
	fmt.Printf("Результат строки абуБАКА: %v\n", uniqueSymbols(test4))
	fmt.Printf("Результат строки aBCD: %v\n", uniqueSymbols(test5))
	fmt.Printf("Результат строки абвгд: %v\n", uniqueSymbols(test6))

}
