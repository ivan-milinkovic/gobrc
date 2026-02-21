package main

import (
	"fmt"
	"time"
)

const file = "external/measurements_10k.txt" // if git does a crlf conversion, it will fail
const buff_size = 50_000                     // ~50kB

// const file = "external/measurements_1M.txt"
// const buff_size = 5_000_000 // ~5MB

// const file = "external/measurements_10M.txt"
// const buff_size = 50_000_000 // ~100MB

// const file = "external/measurements_100M.txt"
// const buff_size = 100_000_000 // ~100MB

// const file = "external/measurements_1B.txt"
// const buff_size = 200_000_000 // ~200MB

func main() {
	var t0 = time.Now()
	// sorted by speed: fastest - slowest
	res, err := process_conc_buff(file, buff_size)
	// res, err := process_seq_scan(file)
	// res, err := process_seq_manual(file)
	// res, err := process_conc_copies(file)
	// res, err := process_conc_slots(file)
	// res, err := process_conc_reads(file)

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
