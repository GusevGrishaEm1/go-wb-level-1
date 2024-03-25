package main

func task11(set1 []int, set2 []int) []int {
	intersection := make(map[int]bool)

	for _, el := range set1 {
		intersection[el] = true
	}

	result := make([]int, 0)

	for _, el := range set2 {
		_, ok := intersection[el]
		if ok {
			result = append(result, el)
		}
	}

	return result
}