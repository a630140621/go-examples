package chapter8

import "testing"

func TestCounterWithoutGoRoutines(t *testing.T) {
	CounterWithoutGoRoutines()
}

func TestCounterWithGoRoutines(t *testing.T)  {
	CounterWithGoRoutines()
}
