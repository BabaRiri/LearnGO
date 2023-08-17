package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data []int
		answer int
	}

	tests := []test {
		{
			[]int{1, 1},
			2,
		},
		{
			[]int{2, 2},
			4,
		},
		{
			[]int{3, 3},
			6,
		},
		{
			[]int{4, 4},
			8,
		},
	}

	for _, v := range tests {
		ans := mySum(v.data...)
		if ans != v.answer {
			t.Error("EXPECTED: ", v.answer," GOT: ", ans)
		}
		
	}
}