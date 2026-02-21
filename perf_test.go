package main

import "testing"

const test_file = file
const test_buff_size = buff_size

func BenchmarkProcessSeqScan(b *testing.B) {
	for range b.N {
		process_seq_scan(test_file)
	}
}

func BenchmarkProcessSeqManual(b *testing.B) {
	for range b.N {
		process_seq_manual(test_file)
	}
}

func BenchmarkProcessConcurrentReads(b *testing.B) {
	for range b.N {
		process_conc_reads(test_file)
	}
}

func BenchmarkProcessConcurrentCopies(b *testing.B) {
	for range b.N {
		process_conc_copies(test_file)
	}
}

func BenchmarkProcessConcurrentSlots(b *testing.B) {
	for range b.N {
		process_conc_slots(test_file)
	}
}

func BenchmarkProcessConcurrentBuffer(b *testing.B) {
	for range b.N {
		process_conc_buff(test_file, test_buff_size)
	}
}
