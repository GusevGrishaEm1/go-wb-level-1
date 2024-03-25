package main

func task12(arr []string) []string {
	mp := make(map[string]struct{})
	for _, el := range arr {
		mp[el] = struct{}{}
	}
	result := make([]string, len(mp))
	i := 0
	for key := range mp {
		result[i] = key
		i++
	}
	return result
}