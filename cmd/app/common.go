package main

import "math/rand"

func generateRandomWord() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	word := make([]byte, 5)
	for i := range word {
		word[i] = letters[rand.Intn(len(letters))]
	}
	return string(word)
}