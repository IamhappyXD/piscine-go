package main

func FindNextPrime(nb int) int {
	if nb <= 1 || nb == 2 {
		return 2
	} else {
		for j := 2; j < nb; j++ {
			if nb%j == 0 {
				return FindNextPrime(nb + 1)
			}
		}
		return nb

	}

}
