package main
import "fmt"

func main(){
	x := []int{1,2,3,4}
	fmt.Println(x)
	 y := []int{4,7,8,9}
	 z := append(x,y...)
	 fmt.Println(z)
	 //deleting elements from slice
	 y = append(x[:2],z[4:]...)
	fmt.Println(y)
}