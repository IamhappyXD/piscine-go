package piscine

func Index(s string, toFind string) int {
	main := []rune(s)
	check := []rune(toFind)
	ind := 0
	count := 0
	total := 0

	for i := range main {
		for j := range check {
			if main[i] == check[j] {
				count++
			} else {
				break
			}
			total++

		}

		if total == count && total > 0 {
			return ind
		}
		count := 0
		total := 0
		ind++
	}
	return -1
}
