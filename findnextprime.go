package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 || nb == 2 {
		return 2
	} else {
		i := nb

		for i < nb+12 {

			if (nb > 10) && ((i%2 == 0) || (i%10 == 2) || (i%10 == 5) || (i%10 == 0)) {
				i++
				continue
			} else {
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
				i++
			}

		}
		return nb

	}

}
