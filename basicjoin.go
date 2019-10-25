package piscine

func BasicJoin(strs []string) string {
	a := ""
	for i := range strs {
		a = a + strs[i]
	}
	return a
}
