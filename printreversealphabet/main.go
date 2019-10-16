package main

import "github.com/01-edu/z01"
import "fmt"

func main() {
	for i := 122; i > 96; i-- {

		ch := string(i)
		fmt.Printf(ch)
	}
	z01.PrintRune(10)
}
