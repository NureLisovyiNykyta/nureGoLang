package fileops

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// CreateAbsIntFile створює файл з модулями цілих чисел.
func CreateAbsIntFile(filename string, count int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	rand.Seed(time.Now().UnixNano())
	writer := bufio.NewWriter(f)
	for i := 0; i < count; i++ {
		num := rand.Intn(201) - 100
		absNum := int(math.Abs(float64(num)))
		_, err := fmt.Fprintln(writer, absNum)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// ReadNumbers читає числа з файлу.
func ReadNumbers(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var numbers []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

// MinOddIndexed знаходить найменше значення серед компонент з непарними номерами.
func MinOddIndexed(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("порожній файл")
	}
	min := math.MaxInt32
	found := false
	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 { // непарні номери (1-based: 1,3,5...)
			if numbers[i] < min {
				min = numbers[i]
			}
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("немає елементів з непарними номерами")
	}
	return min, nil
}