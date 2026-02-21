package main

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func process_conc_consumers(file_path string) (map[string]Stats, error) {
	var file, err = os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var sc = bufio.NewScanner(file)

	var nworkers = runtime.NumCPU()
	var ch_lines = make(chan string, nworkers)          // line input for goroutines
	var ch_res = make(chan map[string]Stats, nworkers)  // goroutines send their results to this channel
	var acc_res = make([]map[string]Stats, 0, nworkers) // accumulated results per worker

	for range nworkers {
		go sub_process_conc_consumers(ch_lines, ch_res)
	}

	for sc.Scan() {
		// too many allocations, string allocation per line
		ch_lines <- sc.Text() // need to send a copy, as Scanner re-uses its internal buffer
	}
	close(ch_lines)

	for res := range ch_res {
		acc_res = append(acc_res, res)
		if len(acc_res) == nworkers {
			break
		}
	}

	res := merge_results(acc_res)

	for k := range res {
		stats := res[k]
		stats.mean /= float64(stats.count)
	}

	return res, nil
}

func sub_process_conc_consumers(ch_lines <-chan string, ch_res chan<- map[string]Stats) {
	var res = make(map[string]Stats)
	for line := range ch_lines {
		parts := strings.Split(line, ";")
		name, temp_str := parts[0], parts[1]
		temp, err := strconv.ParseFloat(temp_str, 64)
		if err != nil {
			panic("Bad input")
		}

		stats, found := res[name]
		if !found {
			stats = Stats{}
			// name2 := string(bname) // escapes as it's added to the map, only escape if needed
			name2 := string(name) // escapes as it's added to the map, only escape if needed
			res[name2] = stats
		}

		stats.min = min(stats.min, temp)
		stats.min = max(stats.max, temp)
		stats.mean += temp // will divide later
		stats.count += 1
	}
	ch_res <- res
}
