package piscine

func NRune(s string, n int) rune {
	check := []rune(s)
	for i := range check {
		if i == n-1 {
			return check[i]
		}
	}
	if n > i || n == i+1 {
		return ""
	}
}
