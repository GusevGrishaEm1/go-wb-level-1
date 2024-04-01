package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func someFunc() (string, error) {
	v := createHugeString(1024) //100
	size := utf8.RuneCountInString(*v)
	if size < 100 {
		return "", errors.New("too short")
	}
	arr1 := (*v)[:100]
	runes := make([]rune, 100)
	copy(runes, []rune(arr1)) 
	return string(runes), nil
}

func createHugeString (size int) *string {
	hugeStr := "i'am a huge string"
	return &hugeStr
}

func task15() {
	justString, err := someFunc()
	if err!= nil {
        fmt.Println(err)
		return
    }
	fmt.Println(justString)
}