package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	} else {
		for i := nb; i < nb+20; i++ {
			check := 0
			if i/10 > 1 && (i%2 == 0 || i%10 == 2 || i%10 == 5 || i%10 == 0) {
				continue
			}

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
