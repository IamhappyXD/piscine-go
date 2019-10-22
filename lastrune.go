package piscine

func LastRune(s string) rune {
	check := []rune(s)
	count := 1
	for i := range check {
		count++
	}
	return check[count]
}
