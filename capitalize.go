package main

import "fmt"

func Capitalize(s string) string {
	check := []rune(s)
	start := 0
	for i := range check {
		if start == 0 {
			if 65 <= check[i] && check[i] <= 90 {
				start++
				continue
			} else if 97 <= check[i] && check[i] <= 122 { //toupper
				check[i] = check[i] - 32
				start++
				continue
			} else {
				start++
			}
		}

		if 65 <= check[i] && check[i] <= 90 { //tolower
			check[i] = check[i] + 32
		} else if (97 <= check[i] && check[i] <= 122) || (check[i] >= 48 && check[i] <= 57) {
			continue
		} else {
			start = 0
		}
	}
	return string(check)
}

func main() {
	fmt.Println(Capitalize("J*+4bGZX(dB}U"))
}
