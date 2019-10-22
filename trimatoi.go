package piscine

func TrimAtoi(s string) int {
	var a []int
	check := []rune(s)
	count := 0
	for i := range check {
		if check[i] >= 48 && check[i] <= 57 {
			a = append(a, int(check[i])-48)
			count++
		}

	}

	if count == 0 {
		return 0
	} else {
		x := 0
		for i := 0; i < count; i++ {
			x *= 10
			x += a[i]

		}

		return x

	}

}
