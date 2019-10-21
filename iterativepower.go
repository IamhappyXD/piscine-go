package main
import "fmt"

func IterativePower(nb int, power int) int {
	result := 1
	for i:=0; i<power;i++{
		result=result*nb
	}
	return result
}

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(IterativePower(arg1, arg2))
}