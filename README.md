# gobrc

A solution to "One Billion Rows Challenge", focusing on performance, without output verification.

Time: 
~13.3s on M1 Pro, NumCPU = 10
~ 9.4s on Ryzen 9800X3D, NumCPU = 16

# Setup

Use the script in external to generate data files.
Scripts come from the official repository: https://github.com/gunnarmorling/1brc/

```sh
cd external
python3 create_measurements.py 1_000_000_000
```

It took 6 minutes 3 seconds to generate measurements_1B.txt file of 14.8 GiB in size.

# Test

```sh
go test -bench . -run NONE -bench=^BenchmarkProcessConcurrentBuffer$ -cpuprofile=cpu.prof
go tool pprof -http=:8080 cpu.prof
```

# Results

1 million lines:
process_conc_buff: 39.846167ms (5MB buffer)
process_seq_scan: 90.34375ms
process_seq_manual: 119.787208ms
process_conc_copies: 470.418333ms
process_conc_slots: 760.252375ms
process_conc_reads: 3340.561875ms

100 million lines:
process_conc_buff: 1.375980583s (100MB buffer)
process_seq_scan: 6.958238958s
process_seq_manual: 9.930498333s
process_conc_copies: 43.551966667s
process_conc_slots: 64.157248208s
process_conc_reads: too long

1 Billion lines:
process_conc_buff: 13.279955208s (200MB buffer)

# Ref

https://pkg.go.dev/io
https://pkg.go.dev/bufio
https://pkg.go.dev/golang.org/x/exp/mmap
