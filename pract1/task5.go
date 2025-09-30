package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task5() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = r.Intn(21) - 10
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%4d ", matrix[i][j])
		}
		fmt.Println()
	}

	sum := 0
	for i := 0; i < 3; i++ {
		if matrix[i][1] > 0 && matrix[i][1]%2 == 0 {
			sum += matrix[i][1]
		}
	}

	fmt.Printf("Sum of positive even numbers in the second column: %d\n", sum)
}