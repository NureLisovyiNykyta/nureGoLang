package main

import (
	"fmt"
	"nureGoLang/lab3/fileops"
)

func task1() {
	filename := "numbers.txt"
	err := fileops.CreateAbsIntFile(filename, 10)
	if err != nil {
		panic(err)
	}
	numbers, err := fileops.ReadNumbers(filename)
	if err != nil {
		panic(err)
	}
	min, err := fileops.MinOddIndexed(numbers)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Найменше значення компонент з непарними номерами: %d\n", min)
}