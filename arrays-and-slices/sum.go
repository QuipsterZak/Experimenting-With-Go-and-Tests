package main

func SumArray(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}
	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, SumSlice(tail))
		}

	}
	return sums
}
