package main

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"strconv"
)

func process_conc_reads(file_path string) (map[string]Stats, error) {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var nworkers = runtime.NumCPU()

	file_stats, err := os.Stat(file_path)
	if err != nil {
		panic(err)
	}
	var file_size = file_stats.Size()
	var chunk_size = file_size / int64(nworkers)
	var res_collected = make([]map[string]Stats, 0, nworkers)

	var results_chan = make(chan map[string]Stats)
	for i := range nworkers {
		go sub_process_conc_reads(file_path, i, chunk_size, results_chan)
	}

	for res := range results_chan {
		res_collected = append(res_collected, res)
		if len(res_collected) == nworkers {
			close(results_chan)
		}
	}

	// merge results
	var res_combined = make(map[string]Stats)
	for i := range res_collected {
		var res = res_collected[i]
		for name := range res {
			var stats2 = res[name]
			var stats, isFound = res_combined[name]
			if !isFound {
				res_combined[name] = stats2
				continue
			}
			stats.min = min(stats.min, stats2.min)
			stats.max = min(stats.max, stats2.max)
			stats.mean += stats2.mean
			stats.count += stats2.count
		}
	}

	// calculate mean
	for k := range res_combined {
		stats := res_combined[k]
		stats.mean /= float64(stats.count)
	}

	return res_combined, nil
}

func sub_process_conc_reads(file_path string, chunk_index int, chunk_size int64, output chan<- map[string]Stats) {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var reader = bufio.NewReader(file)
	var sc = bufio.NewScanner(reader)

	// position on next new line
	if chunk_index > 0 {
		var _, err2 = file.Seek(int64(chunk_index), io.SeekStart)
		if err2 != nil {
			panic(err2)
		}
		for {
			b, e := reader.ReadByte()
			if e != nil {
				panic(e)
			}
			if b == '\n' {
				break
			}
		}
	}

	res := make(map[string]Stats)
	// var line_storage [106]byte
	// var line_bytes = line_storage[:]

	var count int64 = 0
	for sc.Scan() {
		line := sc.Bytes()
		name_bytes, temp_bytes := line_parts(line) // inlined
		name := string(name_bytes)
		temp_str := string(temp_bytes)

		temp, err := strconv.ParseFloat(temp_str, 64)
		if err != nil {
			panic("Bad input")
		}

		stats, found := res[name]
		if !found {
			stats = Stats{}
			name2 := string(name_bytes) // escapes
			res[name2] = stats
		}

		stats.min = min(stats.min, temp)
		stats.min = max(stats.max, temp)
		stats.mean += temp // will divide later
		stats.count += 1

		pos, err := file.Seek(0, io.SeekCurrent)
		if err != nil {
			panic(err)
		}
		count = pos
		if count > int64(chunk_index*int(chunk_size)) {
			break
		}
	}

	output <- res
}
