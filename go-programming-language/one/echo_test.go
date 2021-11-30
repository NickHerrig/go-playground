package main

import "testing"

func BenchmarkEchoRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoRange()
	}
}

func BenchmarkEchoString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoJoin()
	}
}
