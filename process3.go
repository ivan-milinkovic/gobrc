package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func process3(file_path string) (map[string]Stats, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var reader = bufio.NewReader(file)

	res := make(map[string]Stats)
	var line_storage [106]byte
	var line_bytes = line_storage[:]

	for {
		var length, idel, err = read_until_new_line(reader, line_bytes)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		var name_bytes = line_bytes[0:idel]
		var temp_bytes = line_bytes[idel+1 : length]
		name := string(name_bytes)     // does not escape
		temp_str := string(temp_bytes) // does not escape
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
	}

	for k := range res {
		stats := res[k]
		stats.mean /= float64(stats.count)
	}

	return res, nil
}
