package main

import "github.com/01-edu/z01"

func main() {
	
	var aRune string = "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(aRune[i]))
	}
	var bRune rune = '\n'
	z01.PrintRune(bRune)
}
