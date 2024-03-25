package main

import "fmt"

func someFunc() *string {
	v := createHugeString(1024) //100
	if len(*v) < 100 {
		return nil
	}
	return v
}

func createHugeString (size int) *string {
	hugeStr := "i'am a huge string"
	return &hugeStr
}

func task15() {
	var justString *string
	justString = someFunc()
	fmt.Println(*justString)
}