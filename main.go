package main

import (
	"fmt"
	"time"
)

const file = "external/measurements_10k.txt"

func main() {
	var t0 = time.Now()
	res, err := process_seq_scan(file)
	// res, err := process_seq_manual(file)
	// res, err := process_conc_reads(file)
	// res, err := process_conc_consumers(file)

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
