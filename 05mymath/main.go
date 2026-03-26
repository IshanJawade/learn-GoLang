package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcom to math in golang")

	// # random number from math package
	// rand.Seed()  ***Deprecated!!!
	// fmt.Println(rand.Intn(7))

	// # random number from crypto package
	cryptoRandNum, _ := rand.Int(rand.Reader, big.NewInt(7))
	fmt.Println(cryptoRandNum)
}
