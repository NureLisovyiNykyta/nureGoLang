package main

import (
	"fmt"
)

func task2() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	// Перевіряємо, чи трикутник існує
	isTriangle := a+b > c && a+c > b && b+c > a

	// Перевіряємо, чи трикутник різнобічний
	isScalene := isTriangle && a != b && b != c && a != c

	if isScalene {
		fmt.Println("Трикутник різнобічний")
	} else {
		fmt.Println("Трикутник не різнобічний або не існує")
	}
}