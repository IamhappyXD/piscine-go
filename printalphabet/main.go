package main
import "fmt"
func main(){
	for i:= 97;i<122;i++{
	ch:= rune(i)
	fmt.Printf("%c", ch)
	}
	fmt.Printf("\n")
}
