# gobrc

# Results

1 million lines:
process_seq_scan: 90.34375ms
process_seq_manual: 119.787208ms
process_conc_buff: 165.208709ms
process_conc_copies: 470.418333ms
process_conc_slots: 760.252375ms
process_conc_reads: 3340.561875ms

100 million lines:
process_seq_scan: 6.958238958s
process_seq_manual: 9.930498333s
process_conc_buff: 17.151609625s
process_conc_copies: 43.551966667s
process_conc_slots: 64.157248208s
process_conc_reads: too long

# Ref

https://pkg.go.dev/io
https://pkg.go.dev/bufio
https://pkg.go.dev/golang.org/x/exp/mmap
