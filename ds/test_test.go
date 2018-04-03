package main

import (
	"fmt"
	"testing"
)

func Average(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}
	sum := 0.0
	for _, r := range list {
		sum += r
	}
	return sum / float64(len(list))
}

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func main() {
	fmt.Println(Average([]float64{1, 2, 3, 4}))
}
