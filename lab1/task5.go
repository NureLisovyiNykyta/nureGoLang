package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task5() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var matrix [4][5]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = r.Intn(201) - 100
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%4d ", matrix[i][j])
		}
		fmt.Println()
	}

	count := 0
	for j := 0; j < 5; j++ {
		product := 1
		for i := 0; i < 4; i++ {
			product *= matrix[i][j]
		}
		if product > 100 {
			for i := 0; i < 4; i++ {
				if matrix[i][j] < 0 && matrix[i][j]%5 == 0 {
					count++
				}
			}
		}
	}

	fmt.Printf("Number of negative multiples of 5: %d\n", count)
}