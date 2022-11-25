package int

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	exp := 4

	if sum != exp {
		t.Errorf("expected %d got %d", exp, sum)
	}
}

// Example will add example code into documentation and output comment is used to test example
func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// output: 3
}
