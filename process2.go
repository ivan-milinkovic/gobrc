package main

import (
	"bufio"
	"os"
	"strconv"
)

func process2(file_path string) (map[string]Stats, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// var reader io.Reader = file
	var reader = bufio.NewReader(file)

	res := make(map[string]Stats)

	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		line := sc.Bytes()
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
			name2 := string(bname) // escapes as it's added to the map, only escape if needed
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
