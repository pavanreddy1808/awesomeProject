package main
import "fmt"
func main(){
	for  i:=9; i< 30 ;i++{

		fmt.Println("/n/t",i)
		 for j:=0;j<i ;j++{

		 	fmt.Println(j,"\t")
		 }

	}
}
func lp(){
	x:=1
	for {
		if x<20{ break
		}
		fmt.Println(x)

	}
}