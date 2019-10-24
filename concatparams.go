package piscine

func ConcatParams(args []string) string {
	st := ""
	for x, i := range args {
		if x == 0 {
			st = st + i
		} else {
			st = st + "\n" + i
		}

	}
	return st
}
