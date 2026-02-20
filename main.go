package main

import (
	"fmt"
	"time"
)

func main() {
	var t0 = time.Now()
	res, err := process1("test.txt")
	println(time.Since(t0).String())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Result count: %v\n", len(res))
}

type Stats struct {
	mean, min, max float64
	count          int
}
