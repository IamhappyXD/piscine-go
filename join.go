package piscine

func Join(strs []string, sep string) string {

	a := ""
	i := 0
	lenstrs := 0
	for j := range strs {
		j++
		lenstrs = j

	}
	for i < lenstrs {
		a = a + strs[i]
		if i == lenstrs-1 {
			break
		}
		a = a + sep
		i++
	}
	return a
}
