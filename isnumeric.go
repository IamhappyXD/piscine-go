package piscine

func IsNumeric(str string) bool {
	check := []rune(str)
	for i := range check {
		if check[i] >= 48 && check[i] <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
