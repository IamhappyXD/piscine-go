package main

import "fmt"

func TrimAtoi(s string) int {
	var a []int
	check := []rune(s)
	minus := 0
	count := 0
	for i := range check {
		if check[i] >= 48 && check[i] <= 57 {
			a = append(a, int(check[i])-48)
			count++
		}
		if check[i] == 45 {
			minus++
		}
	}

	if count == 0 {
		return 0
	} else {
		x := 0
		for i := 0; i < count; i++ {
			x *= 10
			x += a[i]

		}
		if minus > 0 {
			return 0 - x
		} else {
			return x
		}

	}

}
func main() {
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
}
