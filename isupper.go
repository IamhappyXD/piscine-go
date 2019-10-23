package piscine

func IsUpper(str string) bool {
	check := []rune(str)
	for i := range check {
		if 65 <= check[i] && check[i] <= 90 {
			continue
		} else {
			return false
		}
	}
}
