package chapter8

import "testing"

func TestCounterWithoutGoRoutines(t *testing.T) {
	CounterWithoutGoRoutines()
}

func TestCounterWithGoRoutines(t *testing.T) {
	CounterWithGoRoutines()
}

func BenchmarkCounterWithoutGoRoutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CounterWithoutGoRoutines()
	}
}
// BenchmarkCounterWithoutGoRoutines-4   	   46982	     25354 ns/op	       0 B/op	       0 allocs/op

func BenchmarkCounterWithGoRoutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CounterWithGoRoutines()
	}
}
// BenchmarkCounterWithGoRoutines-4   	    1465	    883588 ns/op	     294 B/op	       4 allocs/op