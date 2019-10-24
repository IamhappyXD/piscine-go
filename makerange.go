package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		var answer []int
		return answer
	} else {
		answer := make([]int, max-min)
		size := max - min
		for i := 0; i < size; i++ {
			answer[i] = min
			min++
		}
		return answer
	}
}
