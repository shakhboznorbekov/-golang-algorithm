package main

import "fmt"

func heapify(arr []int, n, i int) {

	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {

		largest = left
	}

	if right < n && arr[right] > arr[largest] {

		largest = right
	}

	if largest != i {

		arr[i], arr[largest] = arr[largest], arr[i]

		heapify(arr, n, largest)
	}
}

func heapsort(arr []int, n int) {

	for i := n/2 - 1; i >= 0; i-- {

		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {

		arr[0], arr[i] = arr[i], arr[0]

		heapify(arr, i, 0)
	}
}

func main() {

	var arr []int = []int{12, 9, 1, 5, 10, 6}

	n := len(arr)

	heapsort(arr, n)

	for _, i := range arr {

		fmt.Printf("%d ", i)
	}

	fmt.Println()
}