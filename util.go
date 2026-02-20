package main

import (
	"bufio"
	"io"
)

func line_parts(line []byte) ([]byte, []byte) {
	var idelimiter = -1
	for i, c := range line {
		if c == ';' {
			idelimiter = i
			break
		}
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
