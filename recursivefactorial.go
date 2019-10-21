package main
import "fmt"

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb ==0 {
		return 1
	}else{
		return nb * RecursiveFactorial(nb-1)
	}
		
}
func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
}