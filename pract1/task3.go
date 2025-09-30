package main

import (
	"fmt"
	"math"
)

func task3() {
	var element int
	var value float64
	fmt.Scan(&element, &value)

	if value <= 0 {
		fmt.Println("Error: Value must be positive")
		return
	}

	var a, b, h, S float64

	switch element {
	case 1:
		a = value
		b = a * math.Sqrt(2)
		h = a / math.Sqrt(2)
		S = (a * a) / 2

	case 2:
		b = value
		a = b / math.Sqrt(2)
		h = a / math.Sqrt(2)
		S = (a * a) / 2

	case 3:
		h = value
		a = h * math.Sqrt(2)
		b = a * math.Sqrt(2)
		S = (a * a) / 2
		
	case 4:
		S = value
		a = math.Sqrt(2 * S)
		b = a * math.Sqrt(2)
		h = a / math.Sqrt(2)

	default:
		fmt.Println("Error: Invalid element number")
		return
	}

	fmt.Printf("Cathetus: %.2f\n", a)
	fmt.Printf("Hypotenuse: %.2f\n", b)
	fmt.Printf("Height: %.2f\n", h)
	fmt.Printf("Area: %.2f\n", S)
}