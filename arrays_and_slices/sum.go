package main

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i <5; i++ {
		// sum += numbers[i]
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := []int{}

	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}