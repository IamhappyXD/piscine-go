package piscine

func Index(s string, toFind string) int {
	main := []rune(s)
	check := []rune(toFind)
	ind := 0
	count := 0
	total := 0

	for j := range check {
		total++
		if check[j] == 0 {
			return 0
		}
	}

	for i := range main {
		for j := range check {
			if main[i] == check[j] {
				count++
			}

		}
		ind++
	}
	return -1
}
