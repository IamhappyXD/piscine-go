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
}
