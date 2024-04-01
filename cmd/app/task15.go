package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func someFunc() (string, error) {
	v := createHugeString(1024) //100
	if utf8.RuneCountInString(*v) < 100 {
		return "", errors.New("too short")
	}
	return (*v)[:100], nil
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