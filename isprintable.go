package piscine

func IsPrintable(str string) bool {
	check := []rune(str)
	for i := range check {
		if 33 <= check[i] && check[i] <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}
