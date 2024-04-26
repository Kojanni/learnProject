package main

import (
	"fmt"
)

func main() {
	str := "HELOO worLDs"
	newString := ToAlternatingCase(str)
	fmt.Println(str + " - " + newString)
}
