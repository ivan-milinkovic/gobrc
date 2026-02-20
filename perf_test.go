package main

import "testing"

const file_path = "test.txt"

func BenchmarkProcess1(b *testing.B) {
	for range b.N {
		process1(file_path)
	}
}

func BenchmarkProcess2(b *testing.B) {
	for range b.N {
		process2(file_path)
	}
}

func BenchmarkProcess3(b *testing.B) {
	for range b.N {
		process3(file_path)
	}
}
