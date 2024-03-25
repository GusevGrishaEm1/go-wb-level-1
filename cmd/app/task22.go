package main

import (
	"fmt"
	"math/big"
)

func task22() {
	a, _ := new(big.Int).SetString("1355176", 10)
	b, _ := new(big.Int).SetString("3097152", 10)
	
	div := new(big.Int).Div(b, a)
	fmt.Println("Деление:", div)

	mul := new(big.Int).Mul(a, b)
	fmt.Println("Произведение:", mul)

	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма:", sum)

	diff := new(big.Int).Sub(a, b)
	fmt.Println("Разность:", diff)
}