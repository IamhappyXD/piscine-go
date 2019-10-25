package piscine

func SplitWhiteSpaces(str string) []string {
	s := []rune(str)
	count := 0
	for _, i := range s {
		if i == ' ' {
			count++
		}
	}
	count++
	ans := make([]string, count)
	j := 0
	ans[0] = ""

	for _, st := range s {

		if st == ' ' {
			if ans[j] == "" {
				continue
			}

			j++
			ans[j] = ""
			continue
		}
		ans[j] = ans[j] + string(st)

	}
	return ans
}
