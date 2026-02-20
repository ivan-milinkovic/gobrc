package main

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
