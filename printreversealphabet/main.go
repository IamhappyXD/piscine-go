package main

import "fmt"
import "github.com/01-edu/z01"

func main() {
	for i := 122; i > 96; i-- {

		ch := string(i)
		fmt.Printf(ch)
	}
	z01.PrintRune(10)
}
