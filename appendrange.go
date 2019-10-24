package piscine

func AppendRange(min, max int) []int {
	var answer []int
	if max <= min {
		return answer
	} else {
		for min < max {
			answer = append(answer, min)
			min++
		}
		return answer
	}
}
