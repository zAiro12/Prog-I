package main

import (
	"testing"
)

func integra(f func(float64) float64, a, b float64, n int) float64 {
	lung := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		bas1 := f(a + lung*float64(i))
		bas2 := f(a + lung*float64(i+1))
		sum += (bas1 + bas2) * lung / 2
	}
	return sum
}
