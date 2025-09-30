package main

import (
	"fmt"
	"math"
)

func task1() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	closerToA := math.Abs(b-a) < math.Abs(c-a)
	fmt.Println(closerToA)
}