package main

func checkUnique(s string) bool {
	set := make(map[rune]struct{})
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		_, ok := set[r[i]]
		if ok {
			return false
		} else {
			set[r[i]] = struct{}{}
		}
	}
	return true
}