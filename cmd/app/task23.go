package main

func removeElement(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}