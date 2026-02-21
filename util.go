package main

import (
	"bufio"
	"io"
)

const MAX_LINE_LENGTH = 106

func line_parts(line []byte) ([]byte, []byte) {
	var idelimiter = -1
	for i, c := range line {
		if c == ';' {
			idelimiter = i
			break
		}
	}
	if idelimiter == -1 {
		sline := string(line)
		println(sline)
		panic("colon \";\" not found")
	}
	bname := line[0:idelimiter]
	btemp := line[idelimiter+1:]
	return bname, btemp
}

func read_until_new_line(r *bufio.Reader, bytes []byte) (length int, colon_index int, error error) {
	var i = 0
	var ci = -1 // colon index
	for {
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				return 0, -1, err
			}
			panic(err)
		}
		if b == '\n' {
			break
		} else if b == ';' {
			ci = i
		}
		bytes[i] = b
		i++
	}
	return i, ci, nil
}

// Combine array of maps into one map
func merge_results(res_acc []map[string]Stats) map[string]Stats {
	var res_com = make(map[string]Stats) // combined results
	for i := range res_acc {
		var res = res_acc[i]
		for name := range res {
			var stats2 = res[name]
			var stats, isFound = res_com[name]
			if !isFound {
				res_com[name] = stats2
				continue
			}
			stats.min = min(stats.min, stats2.min)
			stats.max = min(stats.max, stats2.max)
			stats.mean += stats2.mean
			stats.count += stats2.count
		}
	}
	return res_com
}
