package main

import "fmt"

// функция возвращает множество уникальных слов,
// полученных из входного множества слов
func setOfUniqueWords(inputSet []string) []string {
	seenWord := make(map[string]bool)               // создаем словарь для отслеживания того, видели ли мы слово или нет
	uniqueWords := make([]string, 0, len(inputSet)) // создаем слайс для хранения уникальных слов

	for _, word := range inputSet {
		if _, ok := seenWord[word]; !ok { // если мы еще не видели такое слово
			seenWord[word] = true                   // устанавливаем флаг в true (слово уже видели)
			uniqueWords = append(uniqueWords, word) // добавляем в слайс это слово, ведь оно уникальное
		}
	}

	return uniqueWords
}

func main() {
	englishSet := []string{"cat", "cat", "dog", "cat", "tree"}
	russianSet := []string{"кошка", "кошка", "собака", "кошка", "дерево"}
	emojiSet := []string{"😺", "😺", "🐶", "😺", "🌳"}
	fmt.Printf("Для множества: %v, уникальными будут слова: %v\n", englishSet, setOfUniqueWords(englishSet))
	fmt.Printf("Для множества: %v, уникальными будут слова: %v\n", russianSet, setOfUniqueWords(russianSet))
	fmt.Printf("Для множества: %v, уникальными будут слова: %v\n", emojiSet, setOfUniqueWords(emojiSet))
}
