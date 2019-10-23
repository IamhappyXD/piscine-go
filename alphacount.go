package piscine

func AlphaCount(str string) int {
	count := 0
	check := []rune(str)
	for i := range check {

		if (65 <= check[i] && check[i] <= 90) || (97 <= check[i] && check[i] <= 122) {
			count++
		}
	}
	return count
}
