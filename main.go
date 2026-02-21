package main

import (
	"fmt"
	"time"
)

// const file = "external/measurements_10k.txt"
// const buff_size = 50_000 // ~50kB for measurements_10k.txt

// const file = "external/measurements_1M.txt"
// const buff_size = 5_000_000 // ~5MB for measurements_1M.txt

// const file = "external/measurements_10M.txt"
// const buff_size = 50_000_000 // ~100MB, for measurements_10M.txt

const file = "external/measurements_100M.txt"
const buff_size = 100_000_000 // ~100MB, for measurements_100M.txt

func main() {
	var t0 = time.Now()
	res, err := process_conc_buff(file, buff_size)
	// res, err := process_seq_scan(file)
	// sorted by speed: fastest - slowest
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
