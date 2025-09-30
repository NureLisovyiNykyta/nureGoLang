package main

import (
	"fmt"
)

// Функція для підрахунку кількості дільників числа
func countDivisors(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}
	return count
}


func task4() {
	var n, m int
	fmt.Print("Введіть два додатні цілі числа n і m (n ≤ m): ")
	fmt.Scan(&n, &m)

	// Перевірка коректності введених даних
	if n > m || n <= 0 {
		fmt.Println("Некоректний відрізок: n має бути меншим або дорівнювати m, і обидва мають бути додатними")
		return
	}

	maxDivisors := 0
	result := n

	// Перебираємо всі числа на відрізку [n, m]
	for i := n; i <= m; i++ {
		divisors := countDivisors(i)
		if divisors > maxDivisors {
			maxDivisors = divisors
			result = i
		}
	}

	fmt.Printf("Число %d має найбільшу кількість дільників: %d\n", result, maxDivisors)
}
