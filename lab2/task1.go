package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функція для введення масиву з консолі
func readArray() []float64 {
	fmt.Println("Введіть елементи масиву через пробіл:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strs := strings.Split(input, " ")

	var arr []float64
	for _, s := range strs {
		num, err := strconv.ParseFloat(s, 64)
		if err == nil {
			arr = append(arr, num)
		}
	}
	return arr
}

// Функція для обробки масиву: знайти перший максимальний елемент і замінити на 0
func processArray(arr []float64) []float64 {
	if len(arr) == 0 {
		return arr
	}

	// Знаходимо максимальне значення
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	// Знаходимо перший індекс з максимальним значенням і замінюємо на 0
	for i, val := range arr {
		if val == maxVal {
			arr[i] = 0
			break
		}
	}
	return arr
}

// Функція для виведення масиву
func printArray(arr []float64) {
	fmt.Print("Оброблений масив: ")
	for _, val := range arr {
		fmt.Printf("%.2f ", val)
	}
	fmt.Println()
}

func task1() {
	arr := readArray()
	arr = processArray(arr)
	printArray(arr)
}