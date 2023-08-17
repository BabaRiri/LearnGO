package main

import "testing"

func TestMySum(t *testing.T) {
	ans := mySum(50, 50)

	if ans != 100 {
		t.Error("EXPECTED: 100, GOT: ", ans)
	}
}