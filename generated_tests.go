package main

import (
	"testing"

	stats "github.com/montanaflynn/stats"
)

func TestCalculateDiffsTest1(t *testing.T) {
	x := []float64{1.0, 2.0, 3.0}
	y := []float64{4.0, 5.0, 6.0}
	meanX, _ := stats.Mean(x)
	meanY, _ := stats.Mean(y)
	diffXY, diffXX, diffYY := CalculateDiffs(x, y, meanX, meanY)
}
func TestCalculateDiffsTest2(t *testing.T) {
	x := []float64{7.0, 8.0, 9.0}
	y := []float64{10.0, 11.0, 12.0}
	meanX, _ := stats.Mean(x)
	meanY, _ := stats.Mean(y)
	diffXY, diffXX, diffYY := CalculateDiffs(x, y, meanX, meanY)
}
