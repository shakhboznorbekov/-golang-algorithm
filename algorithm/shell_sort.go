package main

import "fmt"

func shellSort(slice []int, size int) {

	interval := size / 2

	for interval > 0 {
		for i := interval; i < size; i++ {

			temp := slice[i]
			j := i

			for j >= interval && slice[j-interval] > temp {

				slice[j] = slice[j-interval]
				j -= interval
			}

			slice[j] = temp
		}
		interval /= 2
	}
}

func main() {

	slice := []int{3, 5, 22, 32, 85, 4, 12, 6, 52, 7, 2, 0}

	shellSort(slice, len(slice))

	fmt.Println(slice)
}