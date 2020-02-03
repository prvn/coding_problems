package main

import (
	"sort"
	"testing"
)

func sumInArray(in []int, k int) bool {
	sort.Ints(in) // this gives single loop over array instead of O(n^2)
	start := 0
	end := len(in) - 1
	for start < end {
		if in[start]+in[end] == k {
			return true
		} else if in[start]+in[end] < k {
			start++
		} else {
			end--
		}
	}
	return false
}

func TestSumInArray(t *testing.T) {
	in := []int{10, 15, 3, 7}
	res := sumInArray(in, 17)
	if !res {
		t.Errorf("10+7 = 17, must be true")
	}
}
