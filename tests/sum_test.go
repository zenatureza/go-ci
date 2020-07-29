package tests

import (
	"../lib"
	"testing"
)

func TestSum(t *testing.T) {
	var a float64 = 5
	var b float64 = 5
	total := lib.Sum(a, b)

	if total != 10 {
		t.Error("Sum is wrong")
	}
}
