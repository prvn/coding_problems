package main

import (
	"reflect"
	"testing"
)

func productOfRest(in []int) []int {
	left := make([]int, len(in))
	left[0] = 1

	right := make([]int, len(in))
	right[len(in)-1] = 1

	out := make([]int, len(in))

	for i := 1; i < len(in); i++ {
		left[i] = in[i-1] * left[i-1]
	}

	for i := len(in) - 2; i >= 0; i-- {
		right[i] = in[i+1] * right[i+1]
	}

	for i := 0; i < len(in); i++ {
		out[i] = left[i] * right[i]
	}

	return out
}

func TestProductOfRest(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	expected := []int{120, 60, 40, 30, 24}
	out := productOfRest(in)
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Expected: %v; Received: %v\n", expected, out)
	}
}
