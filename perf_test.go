package main

import "testing"

const test_file = "external/measurements_10k.txt"

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

func BenchmarkProcessConcurrentConsumers(b *testing.B) {
	for range b.N {
		process_conc_cons(test_file)
	}
}

	}
}
