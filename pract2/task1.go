package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Базова структура для демонстрації вбудовування та перевизначення
type BaseArray struct{}

// Базовий метод String(), який перевизначатиметься
func (b BaseArray) String() string {
	return "Базовий масив"
}

// Структура для одновимірного масиву з вбудованою базовою структурою
type Array1D struct {
	BaseArray // Вбудовування для успадкування методів
	data      []int
}

// Перевизначений метод String() для Array1D
func (a Array1D) String() string {
	var sb strings.Builder
	for _, val := range a.data {
		fmt.Fprintf(&sb, "%d ", val)
	}
	return sb.String()
}

// Метод для видалення елементів з парними номерами
func (a *Array1D) RemoveEvenPositions() {
	var newData []int
	for i := 0; i < len(a.data); i++ {
		if (i+1)%2 != 0 { // Зберігаємо непарні позиції
			newData = append(newData, a.data[i])
		}
	}
	a.data = newData
}

// Структура для двовимірного масиву з вбудованою базовою структурою
type Array2D struct {
	BaseArray // Вбудовування для успадкування методів
	data      [][]int
	rows      int
	cols      int
}

// Перевизначений метод String() для Array2D
func (a Array2D) String() string {
	var sb strings.Builder
	for _, row := range a.data {
		for _, val := range row {
			fmt.Fprintf(&sb, "%d\t", val)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// Метод для сортування першого рядка шляхом перестановки стовпців
func (a *Array2D) SortFirstRowByColumns() {
	// Створюємо індекси стовпців
	indices := make([]int, a.cols)
	for i := range indices {
		indices[i] = i
	}

	// Сортуємо індекси на основі значень першого рядка
	sort.SliceStable(indices, func(i, j int) bool {
		return a.data[0][indices[i]] < a.data[0][indices[j]]
	})

	// Переставляємо стовпці відповідно до відсортованих індексів
	for r := 0; r < a.rows; r++ {
		newRow := make([]int, a.cols)
		for c, idx := range indices {
			newRow[c] = a.data[r][idx]
		}
		a.data[r] = newRow
	}
}

// Функція для введення одновимірного масиву з консолі
func InputArray1D() (*Array1D, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введіть елементи одновимірного масиву через пробіл:")
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	strs := strings.Split(input, " ")
	data := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("Помилка: невірне число '%s'", s)
		}
		data[i] = num
	}
	return &Array1D{data: data}, nil
}

// Функція для введення двовимірного масиву з консолі
func InputArray2D() (*Array2D, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введіть кількість рядків матриці:")
	rowStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	rows, err := strconv.Atoi(strings.TrimSpace(rowStr))
	if err != nil {
		return nil, fmt.Errorf("Помилка: невірна кількість рядків")
	}

	fmt.Println("Введіть кількість стовпців матриці:")
	colStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	cols, err := strconv.Atoi(strings.TrimSpace(colStr))
	if err != nil {
		return nil, fmt.Errorf("Помилка: невірна кількість стовпців")
	}

	data := make([][]int, rows)
	for r := 0; r < rows; r++ {
		fmt.Printf("Введіть елементи рядка %d через пробіл:\n", r+1)
		input, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		input = strings.TrimSpace(input)
		strs := strings.Split(input, " ")
		if len(strs) != cols {
			return nil, fmt.Errorf("Помилка: невірна кількість елементів у рядку %d", r+1)
		}
		row := make([]int, cols)
		for c, s := range strs {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("Помилка: невірне число '%s' у рядку %d", s, r+1)
			}
			row[c] = num
		}
		data[r] = row
	}
	return &Array2D{data: data, rows: rows, cols: cols}, nil
}

// Функція для виведення одновимірного масиву (використовує перевизначений String)
func PrintArray1D(a *Array1D) {
	fmt.Print("Одновимірний масив: ")
	fmt.Println(*a) // Виклик перевизначеного String()
}

// Функція для виведення двовимірного масиву (використовує перевизначений String)
func PrintArray2D(a *Array2D) {
	fmt.Println("Двовимірний масив:")
	fmt.Println(*a) // Виклик перевизначеного String()
}

func task1() {
	// Обробка одновимірного масиву
	arr1D, err := InputArray1D()
	if err != nil {
		panic(err)
	}
	fmt.Println("Початковий масив:")
	PrintArray1D(arr1D)
	arr1D.RemoveEvenPositions()
	fmt.Println("Після видалення елементів з парними номерами:")
	PrintArray1D(arr1D)

	// Демонстрація перевизначеного методу для Array1D
	fmt.Println("Демонстрація базового String():", arr1D.BaseArray.String())
	fmt.Println("Демонстрація перевизначеного String():", arr1D.String())

	// Обробка двовимірного масиву
	arr2D, err := InputArray2D()
	if err != nil {
		panic(err)
	}
	fmt.Println("Початкова матриця:")
	PrintArray2D(arr2D)
	arr2D.SortFirstRowByColumns()
	fmt.Println("Після сортування першого рядка перестановкою стовпців:")
	PrintArray2D(arr2D)

	// Демонстрація перевизначеного методу для Array2D
	fmt.Println("Демонстрація базового String():", arr2D.BaseArray.String())
	fmt.Println("Демонстрація перевизначеного String():", arr2D.String())
}