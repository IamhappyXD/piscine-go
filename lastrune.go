package piscine

func LastRune(s string) rune {
	check := []rune(s)
	count := 0
	for i := range check {
		count++
		if check[i] == 0 {
			return 0
		}
	}
	return check[count-1]
}
