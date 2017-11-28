package common

import (
	"time"
	"fmt"
)

type FuncInfo struct {
	Fn   func() error
	Name string
}

type FuncsInfo []FuncInfo

func Benchmark(times int, funcs FuncsInfo) error {
	if times > 1 {
		for _, curFunc := range funcs {
			if err := curFunc.Fn(); err != nil {
				return err
			}
			fmt.Printf("Test run: func \"%s\"", curFunc.Name)
		}
	}
	sumDurations := make([]time.Duration, len(funcs))
	for n := 0; n < times; n++ {
		for ind, curFunc := range funcs {
			start := time.Now()
			if err := curFunc.Fn(); err != nil {
				return err
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
	return nil
}
