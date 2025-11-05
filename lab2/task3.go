package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Інтерфейс для обробки матриці
type MatrixProcessor interface {
	Read() ([][]int, int)
	FindPalindromes(mat [][]int, n int) []int
	Print(pals []int)
}

// Структура, що реалізує інтерфейс
type matrixImpl struct{}

// Функція для введення матриці
func (m *matrixImpl) Read() ([][]int, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введіть порядок n:")
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, _ := strconv.Atoi(nStr)

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Введіть %d рядок через пробіл:\n", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		strs := strings.Split(input, " ")

		mat[i] = make([]int, n)
		for j, s := range strs {
			if j < n {
				num, err := strconv.Atoi(s)
				if err == nil {
					mat[i][j] = num
				}
			}
		}
	}
	return mat, n
}

// Функція для перевірки, чи є число паліндромом
func isPalindrome(num int) bool {
	if num >= 100 {
		return false
	}
	s := strconv.Itoa(num)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Функція для пошуку паліндромів у верхній трикутній області (включаючи діагональ)
func (m *matrixImpl) FindPalindromes(mat [][]int, n int) []int {
	var pals []int
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			num := mat[i][j]
			if isPalindrome(num) {
				pals = append(pals, num)
			}
		}
	}
	return pals
}

// Функція для виведення результату
func (m *matrixImpl) Print(pals []int) {
	fmt.Println("Знайдені числа-паліндроми:")
	for _, p := range pals {
		fmt.Printf("%d ", p)
	}
	fmt.Println()
}

func task3() {
	var proc MatrixProcessor = &matrixImpl{}
	mat, n := proc.Read()
	pals := proc.FindPalindromes(mat, n)
	proc.Print(pals)
}