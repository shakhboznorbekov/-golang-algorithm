package main

import "fmt"

func main() {

	slice := []int{3, 1, 18, 38, 12, 84, 2}

	for i := 0; i < len(slice)-1; i++ {

		min := i

		for j := i + 1; j < len(slice); j++ {

			if slice[j] < slice[min] {

				min = j
			}
		}

		slice[min], slice[i] = slice[i], slice[min]
	}

	fmt.Println("Sorted in Slice: ")
	fmt.Println(slice)

}