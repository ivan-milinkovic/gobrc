package main

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
)

type BuffSlot struct {
	taken atomic.Bool
	buff  [MAX_LINE_LENGTH]byte
	len   int
}

// Reads lines in the main routine, copies them to a free, re-usable buffer slot and sends them to goroutines
func process_conc_slots(file_path string) (map[string]Stats, error) {
	var file, err = os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var sc = bufio.NewScanner(file)

	var nworkers = runtime.NumCPU()
	var ch_lines = make(chan *BuffSlot, nworkers)      // line input for goroutines
	var ch_res = make(chan map[string]Stats, nworkers) // goroutines send their results to this channel
	var acc_res = make([]map[string]Stats, nworkers)   // accumulated results per worker

	var buff_slots = make([]BuffSlot, nworkers)

	for range nworkers {
		go sub_process_conc_slots(ch_lines, ch_res)
	}

	for sc.Scan() {
		var line = sc.Bytes()
		var bs *BuffSlot = nil
		for {
			b := findFreeBuffSlot(&buff_slots)
			if b != nil {
				bs = b
				bs.taken.Store(true)
				break
			}
		}
		bs.taken.Store(true)
		copy(bs.buff[:], line)
		bs.len = len(line)
		ch_lines <- bs
	}
	close(ch_lines)

	var count = 0
	for res := range ch_res {
		acc_res[count] = res
		count++
		if count >= nworkers {
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

func findFreeBuffSlot(buff_slots *[]BuffSlot) *BuffSlot {
	for i := range len(*buff_slots) {
		b := &(*buff_slots)[i]
		if !b.taken.Load() {
			return b
		}
	}
	return nil
}

func sub_process_conc_slots(ch_lines <-chan *BuffSlot, ch_res chan<- map[string]Stats) {
	var res = make(map[string]Stats)
	for buff_slot := range ch_lines {
		var line = buff_slot.buff[:buff_slot.len]
		bname, btemp := line_parts(line) // inlined
		name := string(bname)
		temp_str := string(btemp)
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

		buff_slot.taken.Store(false)
	}
	ch_res <- res
}
