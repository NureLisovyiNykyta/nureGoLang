package main

import (
	"fmt"
	"math"
)

func task3() {
	var start, end, step float64
	fmt.Scan(&start, &end, &step)

	fmt.Println("x\t\ty")
	fmt.Println("-----------------")

	for x := start; x <= end; x += step {
		if x >= 1 && x <= 4 {
			y := math.Pow(math.Sin(2+x), 2) / (2 + x)
			fmt.Printf("%.2f\t\t%.2f\n", x, y)
		} else {
			fmt.Printf("%.2f\t\tundefined\n", x)
		}
	}
}