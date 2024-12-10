package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Переворачиваем слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Соединяем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	// Пример строки
	input := "snooow dog sun"
	// Переворачиваем слова в строке
	result := reverseWords(input)
	// Выводим результат
	fmt.Println(result)
}
