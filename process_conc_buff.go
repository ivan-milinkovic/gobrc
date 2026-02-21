package main

import (
	"io"
	"os"
	"runtime"
	"strconv"
)

// check if it reports the correct number of results as other versions

func process_conc_buff(file_path string, buff_size int) (map[string]Stats, error) {
	var file, err = os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// file_stats, err := os.Stat(file_path)
	// if err != nil {
	// 	panic(err)
	// }
	// var file_size = file_stats.Size()

	var nworkers = runtime.NumCPU()

	var w_output = make(chan map[string]Stats, nworkers) // worker output
	var acc_res = make([]map[string]Stats, 0)

	var w_chunk_size = buff_size / nworkers // chunk size per worker
	var buff = make([]byte, buff_size)

	var iter = 0
	for {
		nread, err := io.ReadAtLeast(file, buff, buff_size)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			panic(err)
		}

		// find last line break
		var last_n = buff_size
		for i := nread - 1; i >= 0; i-- {
			if buff[i] == '\n' {
				last_n = i
				break
			}
		}
		file.Seek(int64(last_n-nread+1), io.SeekCurrent)

		var start = 0
		var end = w_chunk_size
		for range nworkers {
			var i = 0
			end = start + w_chunk_size
			if i == nworkers-1 { // last worker can read through the end which is guaranteed to be a new line by the above logic
				end = last_n + 1
			} else {
				for j := end - 1; j > start; j-- {
					if buff[j] == '\n' {
						end = j + 1
						break
					}
				}
			}

			go sub_process_conc_buff(buff[start:end], w_output)

			start = end
		}

		var res_count = 0
		for res := range w_output {
			acc_res = append(acc_res, res)
			res_count++
			if res_count == nworkers {
				break
			}
		}

		iter++
	}

	var res = merge_results(acc_res)

	for k := range res {
		stats := res[k]
		stats.mean /= float64(stats.count)
	}

	return res, nil
}

func sub_process_conc_buff(buff []byte, output chan<- map[string]Stats) {
	var buff_len = len(buff)
	var res = make(map[string]Stats)
	var lstart = 0
	for {
		var ibreak = find_next_line_break_index(buff, lstart, buff_len)
		var line = buff[lstart:ibreak]

		bname, btemp := line_parts(line)
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

		lstart = ibreak + 1
		if lstart >= buff_len {
			break
		}
	}

	output <- res
}

func find_next_line_break_index(buff []byte, start int, end int) int {
	for i := start; i < end; i++ { // intentionally overshoot
		if buff[i] == '\n' {
			return i
		}
	}
	str := string(buff[start:end])
	println(str)
	panic("new line not found")
}
