package piscine

func NRune(s string, n int) rune {
	check := []rune(s)
	return check[n-1]
}
