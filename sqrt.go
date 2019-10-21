package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return nb
	} else {
		result := 1
		i := 1
		for result <= nb {
			result = i * i
			i++
		}
		if i*i == nb {
			return i - 1
		} else {
			return 0
		}

	}

}
