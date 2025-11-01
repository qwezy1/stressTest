package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func cpuWorker(id int) {
	b, c, g, v := 999999999999999999.0, 999999999999999999.0, 999999999999999999.0, 999999999999999999.0

	var complexCalc func(x float64, depth int) float64
	complexCalc = func(x float64, depth int) float64 {
		if depth == 0 {
			return math.Cos(x) + math.Sqrt(x)
		}
		return complexCalc(math.Sin(x)+math.Pow(x, 1.01), depth-1)
	}

	for {
		g += b + c
		v += complexCalc(v+b, 5)
		b += c
		_ = math.Log(math.Abs(g) + 1.0)
		_ = math.Sqrt(math.Abs(v) + 1.0)
	}
}

func memWorker() {
	mem := make([][][]float64, 0)
	for {
		size := 200 + rand.Intn(800)
		arr := make([][]float64, size)
		for i := range arr {
			arr[i] = make([]float64, size)
			for j := range arr[i] {
				arr[i][j] = float64(i*j) * 0.123456
			}
		}
		mem = append(mem, arr)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// процессор
	for i := 0; i < runtime.NumCPU(); i++ {
		go cpuWorker(i)
	}

	// память
	go memWorker()

	go func() {
		for {
			fmt.Println("Программа работает...")
			time.Sleep(5 * time.Second)
		}
	}()

	select {} // тут блокируется поток
}
