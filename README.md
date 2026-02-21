# gobrc

# Results

1 million lines:
process_seq_scan: 79.7385ms
process_seq_manual: 117.986625ms
process_conc_consumers: 455.19375ms
process_conc_reads: 2986.418083ms

100 million lines:
process_seq_scan: 6.907816583s
process_seq_manual: 9.889747542s
process_conc_consumers: 42.920069208s
process_conc_reads: too long to wait

# Ref

https://pkg.go.dev/io
https://pkg.go.dev/bufio
https://pkg.go.dev/golang.org/x/exp/mmap
