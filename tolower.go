package piscine

func ToUpper(s string) string {
	check := []rune(s)
	for i := range check {
		if 65 <= check[i] && check[i] <= 90 {
			check[i] = check[i] + 32
		}
	}
	return string(check)
}