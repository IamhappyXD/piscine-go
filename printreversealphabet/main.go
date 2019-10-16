package main

import "github.com/01-edu/z01"

func main() {
	var xRune string = "abcdefghijklmnopqrstuvwxyz"
	for i := 25; i >= 0; i-- {
		z01.PrintRune(rune(xRune[i]))
	}
	z01.PrintRune(10)
}
