// Bubble Sort implementation in GoLang
package main

import "fmt"

func main() {

	arr := [10]int{-2, 54, 0, 11, -9, 32, 92, 21, 8, 3}

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted is array: ")
	fmt.Println(arr)
}
