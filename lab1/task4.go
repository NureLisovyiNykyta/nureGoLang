package main

import (
	"fmt"
)

// Удаление элементов с нечётными индексами
func deleteOddIndices(arr []int) []int {
	result := []int{}
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			result = append(result, arr[i])
		}
	}
	return result
}

// Добавление элемента по индексу K
func insertAtK(arr []int, k, value int) []int {
	if k < 0 || k > len(arr) {
		fmt.Println("Некоректний індекс K")
		return arr
	}
	result := append(arr[:k], append([]int{value}, arr[k:]...)...)
	return result
}

// Переворот массива
func reverseArray(arr []int) []int {
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i] = arr[len(arr)-1-i]
	}
	return result
}

// Поиск элемента, равного среднему арифметическому
func findAverageElement(arr []int) (int, bool) {
	if len(arr) == 0 {
		return 0, false
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	avg := float64(sum) / float64(len(arr))
	for _, v := range arr {
		if float64(v) == avg {
			return v, true
		}
	}
	return 0, false
}

func task4() {
	// Массив
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Початковий масив:", arr)

	// 1. Удаление элементов с нечётными индексами
	result1 := deleteOddIndices(arr)
	fmt.Println("Після видалення елементів із непарними індексами:", result1)

	// 2. Добавление элемента (например, значение 10 по индексу 2)
	var value, index int
	fmt.Print("Введіть значення для вставки: ")
	fmt.Scan(&value)
	fmt.Print("Введіть індекс для вставки: ")
	fmt.Scan(&index)
	result2 := insertAtK(arr, index, value)
	fmt.Println("Після вставки", value, "за індексом", index, ":", result2)

	// 3. Переворот масиву
	result3 := reverseArray(arr)
	fmt.Println("Перевернутий масив:", result3)

	// 4. Поиск елемента, рівного середньому арифметичному
	avgElement, found := findAverageElement(arr)
	if found {
		fmt.Println("Елемент, рівний середньому арифметичному:", avgElement)
	} else {
		fmt.Println("Елемент, рівний середньому арифметичному, не знайдено")
	}
}