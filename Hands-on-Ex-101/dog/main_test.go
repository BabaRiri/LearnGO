package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	ans := Years(1)

	if ans != 7 {
		t.Error("Expected: 7 GOT: ", ans)
	}
}

func TestYearsTwo(t *testing.T) {
	ans := YearsTwo(1)

	if ans != 7 {
		t.Error("Expected: 7 GOT: ", ans)
	}
}

func ExampleYears() {
	fmt.Println(Years(1))
	// OUTPUT: 7
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(1))
	// OUTPUT: 7
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(1)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(1)
	}
}