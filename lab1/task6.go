package main

import (
	"fmt"
	"os"
	"strings"
)

// Видаляє повторні входження слів із тексту
func removeDuplicateWords(text string) string {
	// Розбиваємо текст на слова
	words := strings.Fields(text)
	// Мапа для відстеження унікальних слів
	seen := make(map[string]bool)
	// Слайс для збереження унікальних слів у порядку появи
	result := []string{}

	// Перебираємо слова
	for _, word := range words {
		// Якщо слово ще не бачили, додаємо до результату
		if !seen[word] {
			seen[word] = true
			result = append(result, word)
		}
	}

	// Об'єднуємо слова назад у рядок
	return strings.Join(result, " ")
}

func task6() {
	// Зчитуємо текст із файлу input.txt
	data, err := os.ReadFile("task6_input.txt")
	if err != nil {
		fmt.Println("Помилка при зчитуванні файлу:", err)
		return
	}

	// Перетворюємо байти в рядок
	text := string(data)

	// Видаляємо повторні слова
	result := removeDuplicateWords(text)

	// Виводимо результат
	fmt.Println("Текст після видалення повторних слів:", result)
}