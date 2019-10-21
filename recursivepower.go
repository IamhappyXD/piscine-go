package main

import (
	"fmt"
)

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(RecursivePower(arg1, arg2))
}
