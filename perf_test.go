package main

import "testing"

const file_path = "test.txt"

func BenchmarkProcess0(b *testing.B) {
	for range b.N {
		process1(file_path)
	}
}
