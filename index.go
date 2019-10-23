<<<<<<< HEAD
package main
import "fmt"
func Index(s string, toFind string) int {
	main := []rune(s)
	check := []rune(toFind)
	chlen :=0
	mlen :=0
	for j := range check {
		if check[j] == 0 {
			return 0
		}
		chlen++
	}
	for j := range main {
		if main[j] == 0 {
			return 0
		}
		mlen++
	}
	if chlen<= mlen{
		j :=0
		ind :=0
		for i :=0;i<mlen;i++{
			if(main[i]==check[j]){
				j++
			}else{
				ind=i
				j=0
			}
			if j == chlen{	
				break
			}
		}
		if(j == chlen){
			return ind+1
		}else{
			return -1
		}
	
	}else{
		return -1
	}

}
func  main(){
	fmt.Println(Index("Amina", "Ami"))
=======
package piscine

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

	for i := range main {
		for j := range check {
			if main[i] == check[j] {
				count++
			}

		}

		if total == count {
			return ind
		} else {
			ind++
			count = 0
			total = 0
		}

	}
	return -1
>>>>>>> 0c77c8c142162b62eea3b070c6ec68f2d5801038
}
