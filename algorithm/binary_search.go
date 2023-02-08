package main 

import "fmt"

func main() {

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	search := 5
	flag := false

	low := 0
	high := len(slice) - 1

	for low <= high {

		mid := (low + high) / 2
		if search > slice[mid] {

			low = mid + 1
		} else if search < slice[mid] {

			high = mid - 1
		} else {

			flag = true
			break
		}
	}

	if !flag {

		fmt.Println("Son topilmadi!! ")
	} else {

		fmt.Println("Son toplidi ", search)
	}
}