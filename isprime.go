package piscine

func IsPrime(nb int) bool {
	if 1 >= nb {
		return false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
			if nb%(i*i) == 0 {
				return false
			}
		}
		return true
	}

}
