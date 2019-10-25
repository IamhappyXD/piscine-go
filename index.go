package main
import "fmt"

func Index(s string, toFind string) int {
	main := []rune(s)
	check := []rune(toFind)
	ind := 0
	count := 0
	total := 0

	for j := range check {
		total++
		if check[j] == 0 {
			return 0
		}
	}
	for j := range main {
		count++
		if main[j] == 0 {
			return 0
		}
	}
	if total>count{
		return -1
	}else{
		equal :=0
		for i := range main {
			for j :=i; j<total; j++ {
				if check[j] == main[i] {
					equal++
					
				}
	
			}
			fmt.Println(equal)
			if equal ==total{
				return ind
			}
			equal =0
			ind++
		}
		return -1
	}



}

func main(){
	fmt.Println(Index("8eLp47Gz~5{/c", "eLp4"))
}