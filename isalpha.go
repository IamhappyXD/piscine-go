package piscine

func IsAlpha(str string) bool {
	check := []rune(str)
	for i := range check {
		if (65 <= check[i] && check[i] <= 90) || 97 <= check[i] && check[i] <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
