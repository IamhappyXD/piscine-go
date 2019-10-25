package piscine

func ToUpper(s string) string {
	check := []rune(s)
	for i := range check {
		if 97 <= check[i] && check[i] <= 122 {
			check[i] = check[i] - 32
		}
	}
	return string(check)
}
