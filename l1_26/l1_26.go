package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		test1 = "abcd"
		test2 = "abCdefAaf"
		test3 = "aabcd"
		test4 = "абуБАКА"
		test5 = "aBCD"
	)
	fmt.Printf("Результат строки abcd: %v\n", uniqueSymbols(test1))
	fmt.Printf("Результат строки abCdefAaf: %v\n", uniqueSymbols(test2))
	fmt.Printf("Результат строки aabcd: %v\n", uniqueSymbols(test3))
	fmt.Printf("Результат строки абуБАКА: %v\n", uniqueSymbols(test4))
	fmt.Printf("Результат строки abcd: %v\n", uniqueSymbols(test5))

}

// функция для проверки на то, что символы в строке уникальны (регистр не важен)
// пример: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false (повторяются a\a)
func uniqueSymbols(inputString string) bool {
	seen := make(map[rune]bool)                  // создаем словарь, где ключ - это символ, а значение - встречался ли такой символ
	lowerRegistr := strings.ToLower(inputString) // переводим все символы в один регистр (нижний)

	for _, symbol := range lowerRegistr { // итерируемся по символам строки в нижнем регистре
		if seen[symbol] { // если символ уже true, значит символ не уникальный и можно вернуть false
			return false
		}
		seen[symbol] = true // отмечаем символ как встреченный
	}

	return true // возвращаем true, если все символы уникальны
}
