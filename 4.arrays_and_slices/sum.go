package main

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i <5; i++ {
		// sum += numbers[i]
	for _, num := range numbers {
		// fmt.Println(num)
		sum += num
	}
	return sum
}

//variadic args ...[]int
func SumAll(slices ...[]int) []int {
	sums := []int{}

	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int{
	result := []int{}
	for _, slice := range slices {

		if len(slice) == 0 {
			result = append(result, 0)
		} else {
			tail := slice[1:]
			result = append(result, Sum(tail))
		}
	}
	return result
}