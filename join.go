package piscine

func Join(strs []string, sep string) string {

	a := ""
	for i := range strs {
		a = a + strs[i] + sep
	}
	return a
}
