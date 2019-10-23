package piscine

func IsLower(str string) bool {
	check := []rune(str)
	for i := range check {
		if 97 <= check[i] && check[i] <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
