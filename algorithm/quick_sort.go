package main

import (
	"fmt"
)

func quicksort(slice []int) []int {

	if len(slice) <= 1 {
		return slice
	}

	median := slice[len(slice)/2]

	low_part := make([]int, 0, len(slice))
	high_part := make([]int, 0, len(slice))
	middle_part := make([]int, 0, len(slice))

	for _, item := range slice {

		if item < median {
			low_part = append(low_part, item)
		} else if item > median {
			high_part = append(high_part, item)
		} else if item == median {
			middle_part = append(middle_part, item)
		}
	}

	low_part = quicksort(low_part)
	high_part = quicksort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part

}

func main() {

	slice := []int{8, 2, 0, 1, -1, 6, 7, 9}

	slice = quicksort(slice)

	fmt.Println(slice)
}