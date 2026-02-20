package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func process1(file_path string) (map[string]Stats, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	res := make(map[string]Stats)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, ";")
		name := parts[0]
		temp_str := parts[1]
		temp, err := strconv.ParseFloat(temp_str, 64)
		if err != nil {
			return nil, err
		}

		stats, found := res[name]
		if !found {
			stats = Stats{}
			res[name] = stats
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
