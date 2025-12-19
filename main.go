package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func slowFunction() {
	total := 0
	for i := 0; i < 1_000_000_00; i++ {
		total += i
	}
	fmt.Println("Total:", total)
}

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	start := time.Now()
	slowFunction()
	fmt.Println("Execution time:", time.Since(start))
}
