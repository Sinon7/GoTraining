package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

func main() {
	// Example usage
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(arr)
	fmt.Println("sort result arr:", arr)
}
