package main

import "strings"

func reverseWords(s string) string {
    words := strings.Fields(s)
	length := len(words)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
    return strings.Join(words, " ")
}