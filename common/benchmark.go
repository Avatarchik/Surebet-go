package common

import (
	"time"
	"fmt"
)

type TestFunc func() error

type FuncInfo struct {
	Fn   TestFunc
	Name string
}

func Benchmark(times int, funcs []FuncInfo) {
	sumDurations := make([]time.Duration, len(funcs))
	for n := 0; n < times; n++ {
		for ind, curFunc := range funcs {
			start := time.Now()
			if err := curFunc.Fn(); err != nil {
				panic(err)
			}
			duration := time.Since(start)
			sumDurations[ind] += duration
			fmt.Printf("Func \"%s\" time: %s\n", curFunc.Name, duration)
		}
		fmt.Println()
	}
	for ind, curFunc := range funcs {
		fmt.Printf("\nFunc \"%s\" average time: %s\n", curFunc.Name, sumDurations[ind]/time.Duration(times))
	}
}
