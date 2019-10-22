package piscine

func TrimAtoi(s string) int {
	var a []int
	check := []rune(s)
	minus := 0
	count := 0
	for i := range check {
		if check[i] >= 48 && check[i] <= 57 {
			a = append(a, int(check[i])-48)
			count++
		}
		if check[i] == 45 {
			minus++
		}
	}

	if count == 0 {
		return 0
	} else {
		x := 0
		for i := 0; i < len(a); i++ {
			x *= 10
			x += a[i]

		}
		if minus > 0 {
			return -x
		} else {
			return x
		}

	}

}
