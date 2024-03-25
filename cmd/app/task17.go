package main

func binarySearch(arr []int, val int) bool {
	quickSort(arr)
	left := 0
	right := len(arr) - 1
	for left <= right {
		center := left + (right-left)/2

		if arr[center] == val {
			return true
		}

		if arr[center] > val {
			right = center - 1
		} else {
			left = center + 1
		}
	}
	return false
}