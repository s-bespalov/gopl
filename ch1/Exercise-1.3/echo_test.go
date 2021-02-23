package main

import "testing"

func BenchmarkEchoJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoJoin()
	}
}

func BenchmarkEchoFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoFor()
	}
}
