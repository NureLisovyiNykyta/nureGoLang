package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Інтерфейс для обробки рядків
type StringProcessor interface {
	Process(text string) string
}

// Структура, яка реалізує інтерфейс
type VowelRemover struct{}

// Метод структури — реалізація інтерфейсу
func (vr VowelRemover) Process(text string) string {
	words := strings.Fields(text)
	var result []string

	for _, word := range words {
		runes := []rune(word)
		last := runes[len(runes)-1]

		if !isVowel(unicode.ToLower(last)) {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

// Допоміжна функція: перевірка голосної
func isVowel(r rune) bool {
	vowels := "аеєиіїоуюяaeiouy"
	return strings.ContainsRune(vowels, r)
}

// Окрема функція, що приймає будь-який тип інтерфейсу StringProcessor
func ProcessString(sp StringProcessor, text string) string {
	return sp.Process(text)
}

func task4() {
	text := "Мама мила раму і вікно та стіл"
	fmt.Println("Початковий рядок:", text)

	remover := VowelRemover{}
	result := ProcessString(remover, text)

	fmt.Println("Результат:", result)
}
