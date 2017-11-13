package main

import "testing"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func TestSum(t *testing.T) {
	if sum() != 0 {
		t.Errorf(`sum() = %#v, want %#v`, sum(), 0)
	}

	if sum(1) != 1 {
		t.Errorf(`sum(1) = %#v, want %#v`, sum(1), 1)
	}

	if sum(1, 2, 3) != 6 {
		t.Errorf(`sum(1,2,3) = %#v, want %#v`, sum(1, 2, 3), 6)
	}

}

func TestSumWithExpand(t *testing.T) {
	values := []int{1, 2, 3, 4}

	if sum(values...) != 10 {
		t.Errorf(`sum(values) = %#v, want %#v`, sum(values...), 10)
	}
}
