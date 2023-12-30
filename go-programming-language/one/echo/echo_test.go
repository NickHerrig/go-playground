// use go test -bench=. to run the benchmark
package main

import "testing"

func BenchmarkOne(b *testing.B){
	for i := 0; i < b.N; i++ {
		one()
	}
}

func BenchmarkTwo(b *testing.B){
	for i := 0; i < b.N; i++ {
		two()
	}
}