package main

import (
	"fmt"
	"strings"
)
// Інтерфейс для аналізу рядків (друга частина)
type StringAnalyzer interface {
	Analyze(text string) (int, int)
}

// Структура, яка реалізує інтерфейс для другої частини
type WordAnalyzer struct{}

// Метод структури — реалізація інтерфейсу
func (wa WordAnalyzer) Analyze(text string) (count int, minLen int) {
	words := strings.Fields(text)
	minLen = -1
	for _, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			if minLen == -1 || len(runes) < minLen {
				minLen = len(runes)
			}
			if runes[0] == runes[len(runes)-1] {
				count++
			}
		}
	}
	if minLen == -1 {
		minLen = 0
	}
	return
}

// Окрема функція, що приймає будь-який тип інтерфейсу StringAnalyzer
func AnalyzeString(sa StringAnalyzer, text string) (int, int) {
	return sa.Analyze(text)
}

func task5() {
	text := "Мама мила раму і вікно та силабус"
	fmt.Println("Початковий рядок:", text)

	// Друга частина
	analyzer := WordAnalyzer{}
	count, minLen := AnalyzeString(analyzer, text)
	fmt.Println("Кількість слів з однаковими першим і останнім символами:", count)
	fmt.Println("Довжина найкоротшого слова:", minLen)
}