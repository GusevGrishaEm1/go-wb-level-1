package main

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	inx := partition(arr)
	quickSort(arr[:inx])
	quickSort(arr[inx+1:])
}

func partition(arr []int) int {
	end := len(arr) - 1
	pivot := arr[end]
	i := 0
	for j := range arr {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}