package piscine

func Join(strs []string, sep string) string {

	a := ""
	i := 0
	for i < len(strs) {
		a = a + strs[i]
		if i == len(strs)-1 {
			break
		}
		a = a + sep
		i++
	}
	return a
}
