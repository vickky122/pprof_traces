package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func workA() {
	sum := 0
	for i := 0; i < 50_000_000; i++ {
		sum += i
	}
}

func workB() {
	sum := 0
	for i := 0; i < 100_000_000; i++ {
		sum += i
	}
}

func orchestrator() {
	workA()
	workB()
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
	orchestrator()
	fmt.Println("Execution time:", time.Since(start))
}
