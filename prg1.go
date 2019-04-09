package main
import "fmt"
// calling function with variable and currect them
func main (){
	fmt.Println("any error ", true)
	x := 42 ; y:=3
	cla(x,y)
	fmt.Println(x, y)
	fmt.Println(x*y)

}
// global declartion of variable
var a,b int
// delcare local variable 
func cla(int, int){

	fmt.Println(a*b)
}