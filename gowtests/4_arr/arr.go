package arr

// array has fixed items while slice does not
func Sum(numbers []int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, arr := range numbersToSum {
		sum = append(sum, Sum(arr))
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	for _, arr := range numbersToSum {
		var tail int
		if len(arr) < 1 {
			tail = 0
		} else {
			tail = Sum(arr[1:])
		}
		sum = append(sum, tail)
	}
	return sum
}

func main() {
}
