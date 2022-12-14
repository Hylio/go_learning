package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	if sum != want {
		t.Errorf("got '%d' but want '%d'", sum, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 2)
	}
}

func BenchmarkShowSeqs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShowSeqs(100)
	}
}
