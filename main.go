package main

import (
	"math"
	"runtime"
)

func main() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go cpuWorker()
	}

	go memWorker()

	select {}
}

func cpuWorker() {
	var b, c, g, v float64
	b, c, g, v = 999999999999999999, 999999999999999999, 999999999999999999, 999999999999999999

	for {

		g = g + b + c
		g = math.Sqrt(g)
		v = math.Cos(v)
		b += c

		_ = math.Pow(v, 2.71828)
		_ = math.Sin(g)
	}
}

func memWorker() {
	mem := make([][]float64, 0)
	for {
		arr := make([]float64, 1000000)
		for i := range arr {
			arr[i] = float64(i) * 0.123456
		}
		mem = append(mem, arr)
	}
}
