package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		answer := make([]int, 0)
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
