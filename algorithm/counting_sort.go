// Counting Sorting implemenetation in GoLang

package main

import (
	"fmt"
)

func countSort(slice []int, size int) {

	max := slice[0]

	for _, i := range slice {

		if i > max {
			max = i
		}
	}

	var count []int

	for i := 0; i <= max; i++ {

		count = append(count, 0)
	}

	for _, i := range slice {

		count[i]++
	}

	for i := 1; i <= max; i++ {

		count[i] += count[i-1]
	}

	var output []int

	for i := 0; i <= max; i++ {
		output = append(output, 0)
	}

	for i := size - 1; i >= 0; i-- {

		output[count[slice[i]]-1] = slice[i]
		count[slice[i]]--
	}

	for i := 0; i < size; i++ {

		slice[i] = output[i]
	}

}

func main() {

	slice := []int{4, 2, 2, 3, 5, 4, 8, 7, 9}
	size := len(slice)

	countSort(slice, size)

	fmt.Println("Sorted is slice: ")
	for _, i := range slice {

		fmt.Printf("%d ", i)
	}
	fmt.Println()
}