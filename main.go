package main

import "fmt"

func computeNumber(numbers []int, targetNumber int) []int {
	// i use for return index position
	// n is current number
	for i, n := range numbers {

		// j use for return index position to plus with another number
		for j, m := range numbers {
			if n+m == targetNumber && i != j {
				return []int{i, j}
			}

		}
	}
	return []int{}
}

func main() {

	listNumber := []int{3, 2}
	sum := 3

	result := computeNumber(listNumber, sum)

	fmt.Println(result)

}
