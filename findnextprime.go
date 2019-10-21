package piscine

func FindNextPrime(nb int) int {
	if 1 >= nb {
		return 2
	} else {
		for i := nb; i < nb*2; i++ {
			check := 0
			for j := 2; j < i; j++ {
				if i%j == 0 {
					check++
					break
				}
				if i%(j*j) == 0 {
					check++
					break
				}
			}
			if check == 0 {
				return i
			}
		}
		return nb
	}

}
