package piscine

func Capitalize(s string) string {
	check := []rune(s)
	start := 0
	for i := range check {
		if start == 0 {
			if 97 <= check[i] && check[i] <= 122 { //toupper
				check[i] = check[i] - 32
				start++
			} else if 65 <= check[i] && check[i] <= 90 {
				start++
			} else {
				continue
			}

			continue
		}
		if 65 <= check[i] && check[i] <= 90 { //tolower
			check[i] = check[i] + 32
		} else if 97 <= check[i] && check[i] <= 122 {
			continue
		} else {
			start = 0
		}
	}
	return string(check)
}
