package piscine

func main() {
	count := 0
	for x := range os.Args {
		count = x
	}
	i := 1
	for i <= count {
		out := []rune(os.Args[i])
		for _, j := range out {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
		i++
	}

}
