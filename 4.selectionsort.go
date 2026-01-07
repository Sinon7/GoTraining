package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	// Example usage
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(arr)
	fmt.Println("sort result arr:", arr)
}
