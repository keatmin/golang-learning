package iter

import "fmt"

func Repeat(rep string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += rep
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("a"))
}
