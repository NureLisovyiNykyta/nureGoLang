package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sync"
)

type Interval struct {
	A, B float64
}

type RootResult struct {
	Method     string
	Root       float64
	Iterations int
	Accuracy   float64
	Err        error
}

var f = func(x float64) float64 {
	return x*x*x - 6*x + 2
}

var df = func(x float64) float64 {
	return 3*x*x - 6
}

func scanPart(start, end, h float64, ch chan<- []Interval, wg *sync.WaitGroup) {
	defer wg.Done()
	var locs []Interval
	x := start
	for x < end {
		y := math.Min(x+h, end)
		fx, fy := f(x), f(y)
		if fx*fy <= 0 {
			locs = append(locs, Interval{x, y})
		}
		x = y
	}
	if len(locs) > 0 {
		ch <- locs
	}
}

func golden(a, b float64, eps float64) (float64, int, error) {
	if f(a)*f(b) > 0 {
		return 0, 0, errors.New("f(a) та f(b) одного знаку")
	}
	iters := 0
	invPhi := (math.Sqrt(5) - 1) / 2
	for b-a > eps {
		x1 := a + (1-invPhi)*(b-a)
		x2 := a + invPhi*(b-a)
		f1, f2 := f(x1), f(x2)
		iters++
		if f(a)*f1 <= 0 {
			b = x1
		} else if f2*f(b) <= 0 {
			a = x2
		} else {
			a = x1
			b = x2
		}
		if iters > 100 {
			return (a + b) / 2, iters, errors.New("занадто багато ітерацій")
		}
	}
	return (a + b) / 2, iters, nil
}

func newton(x0, eps float64) (float64, int, error) {
	x := x0
	iters := 0
	for iters < 100 {
		fx := f(x)
		dfx := df(x)
		if math.Abs(dfx) < 1e-12 {
			panic("похідна занадто мала")
		}
		delta := fx / dfx
		x -= delta
		iters++
		if math.Abs(delta) < eps {
			return x, iters, nil
		}
	}
	return x, iters, errors.New("не сходить")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func argmin(a, b int) int {
	if a < b {
		return 0
	}
	return 1
}

func minf(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func maxf(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func argminf(a, b float64) int {
	if a < b {
		return 0
	}
	return 1
}

func task3() {
	fmt.Println("Програма для наближеного розв'язання рівняння x³ - 6x + 2 = 0")
	a, b := 2.0, 3.0
	h, eps := 0.01, 1e-3
	numParts := 10
	partSize := (b - a) / float64(numParts)
	chScan := make(chan []Interval, numParts)
	var wgScan sync.WaitGroup
	for i := 0; i < numParts; i++ {
		wgScan.Add(1)
		start := a + float64(i)*partSize
		end := a + float64(i+1)*partSize
		go scanPart(start, end, h, chScan, &wgScan)
	}
	go func() {
		wgScan.Wait()
		close(chScan)
	}()
	var locs []Interval
	for parts := range chScan {
		locs = append(locs, parts...)
	}
	if len(locs) == 0 {
		fmt.Println("Корені не знайдено")
		return
	}
	fmt.Printf("Знайдено локалізаційний інтервал: [%.3f, %.3f]\n", locs[0].A, locs[0].B)
	interval := locs[0]
	chRes := make(chan RootResult, 2)
	var wgRes sync.WaitGroup
	// Золотий перетин
	wgRes.Add(1)
	go func() {
		defer wgRes.Done()
		root, iters, err := golden(interval.A, interval.B, eps)
		acc := math.Abs(f(root))
		chRes <- RootResult{"Золотий перетин", root, iters, acc, err}
	}()
	// Ньютон
	wgRes.Add(1)
	go func() {
		defer wgRes.Done()
		var result RootResult
		func() {
			defer func() {
				if r := recover(); r != nil {
					result = RootResult{"Ньютон", 0, 0, 0, fmt.Errorf("відновлено з паніки: %v", r)}
				}
			}()
			x0 := (interval.A + interval.B) / 2
			root, iters, err := newton(x0, eps)
			acc := math.Abs(f(root))
			result = RootResult{"Ньютон", root, iters, acc, err}
		}()
		chRes <- result
	}()
	go func() {
		wgRes.Wait()
		close(chRes)
	}()
	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Printf("Помилка створення файлу: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Fprintln(file, "Метод\tКорінь\tІтерацій\tТочність\tПомилка")
	var results []RootResult
	for res := range chRes {
		results = append(results, res)
		if res.Err != nil {
			fmt.Printf("%s: помилка - %v\n", res.Method, res.Err)
		} else {
			fmt.Printf("%s: корінь = %.6f, ітерацій = %d, точність = %.2e\n", res.Method, res.Root, res.Iterations, res.Accuracy)
		}
		errStr := ""
		if res.Err != nil {
			errStr = res.Err.Error()
		}
		fmt.Fprintf(file, "%s\t%.6f\t%d\t%.2e\t%s\n", res.Method, res.Root, res.Iterations, res.Accuracy, errStr)
	}
	// Порівняння
	if len(results) == 2 {
		r1 := results[0]
		r2 := results[1]
		if r1.Err == nil && r2.Err == nil {
			diff := math.Abs(r1.Root - r2.Root)
			fmt.Printf("Різниця між коренями методів: %.6f\n", diff)
			fmt.Fprintf(file, "Різниця між коренями: %.6f\n", diff)
			fmt.Printf("Метод з меншою кількістю ітерацій: %s (%d vs %d)\n",
				map[int]string{0: r1.Method, 1: r2.Method}[argmin(r1.Iterations, r2.Iterations)],
				min(r1.Iterations, r2.Iterations), max(r1.Iterations, r2.Iterations))
			fmt.Fprintf(file, "Метод з меншою кількістю ітерацій: %s (%d ітерацій)\n",
				map[int]string{0: r1.Method, 1: r2.Method}[argmin(r1.Iterations, r2.Iterations)],
				min(r1.Iterations, r2.Iterations))
			fmt.Printf("Метод з вищою точністю: %s (%.2e vs %.2e)\n",
				map[int]string{0: r1.Method, 1: r2.Method}[argminf(r1.Accuracy, r2.Accuracy)],
				minf(r1.Accuracy, r2.Accuracy), maxf(r1.Accuracy, r2.Accuracy))
			fmt.Fprintf(file, "Метод з вищою точністю: %s (%.2e)\n",
				map[int]string{0: r1.Method, 1: r2.Method}[argminf(r1.Accuracy, r2.Accuracy)],
				minf(r1.Accuracy, r2.Accuracy))
		}
	}
	fmt.Println("Результати збережено в results.txt")
}