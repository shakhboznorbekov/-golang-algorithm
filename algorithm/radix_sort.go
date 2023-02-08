

// Radix sort in Golang

package main

import "fmt"

func getMax(arr []int, n int) int {

	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)

	return max
}

func countingSort(arr []int, size, place int) {

	var output [10]int
	max := (arr[0] / place) % 10

	for i := 1; i < size; i++ {
		if ((arr[i] / place) % 10) > max {
			max = arr[i]
		}
	}

	var count []int

	for i := 0; i < max; i++ {
		count = append(count, 0)
	}

	for i := 0; i < size; i++ {
		count[(arr[i]/place)%10]++
	}
	for i := 1; i < size+2; i++ {
		count[i] += count[i-1]
	}

	for i := size - 1; i >= 0; i-- {
		output[count[(arr[i]/place)%10]-1] = arr[i]
		count[(arr[i]/place)%10]--
	}

	for i := 0; i < size; i++ {
		arr[i] = output[i]
	}
}

func radixsort(arr []int, size int) {

	max := getMax(arr, size)

	for place := 1; max/place > 0; place *= 10 {
		countingSort(arr, size, place)
	}
}

func main() {

	arr := []int{121, 432, 564, 23, 1, 45, 788}
	n := len(arr)
	radixsort(arr, n)

	for _, i := range arr {

		fmt.Printf("%d ", i)
	}
	fmt.Println()
}