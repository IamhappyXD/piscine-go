package piscine

func Compare(a, b string) int {
	first := []rune(a)
	second := []rune(b)
	alen := 0
	blen := 0
	for j := range first {
		if first[j] == 0 {
			return 1
		}
		alen++
	}

	for j := range second {
		if second[j] == 0 {

			return -1
		}
		blen++
	}

	i := 0
	for i < alen && i < blen {
		if first[i] > second[i] {

			return 1
		} else if first[i] < second[i] {
			return -1
		} else {
			i++
		}
	}
	if alen > blen {
		return 1
	} else if alen == blen {
		return 0
	} else {
		return -1
	}

}
