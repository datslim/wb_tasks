package main

import "fmt"

// функция возвращает множество уникальных слов,
// полученных из входного множества слов
func sequenceOfUniqueWords(inputSequence []string) []string {
	seenWord := make(map[string]bool)                    // создаем словарь для отслеживания того, видели ли мы слово или нет
	uniqueWords := make([]string, 0, len(inputSequence)) // создаем слайс для хранения уникальных слов

	for _, word := range inputSequence {
		if _, ok := seenWord[word]; !ok { // если мы еще не видели такое слово
			seenWord[word] = true                   // устанавливаем флаг в true (слово уже видели)
			uniqueWords = append(uniqueWords, word) // добавляем в слайс это слово, ведь оно уникальное
		}
	}

	return uniqueWords
}

func main() {
	englishSeq := []string{"cat", "cat", "dog", "cat", "tree"}
	russianSeq := []string{"кошка", "кошка", "собака", "кошка", "дерево"}
	emojiSeq := []string{"😺", "😺", "🐶", "😺", "🌳"}
	fmt.Printf("Для множества: %v, уникальными будут слова: %v\n", englishSeq, sequenceOfUniqueWords(englishSeq))
	fmt.Printf("Для множества: %v, уникальными будут слова: %v\n", russianSeq, sequenceOfUniqueWords(russianSeq))
	fmt.Printf("Для множества: %v, уникальными будут слова: %v\n", emojiSeq, sequenceOfUniqueWords(emojiSeq))
}
