package main
import "fmt"
const(
	a = 2
	b int = 3
)
func main(){
	fmt.Println("This is my second hands on test part1")
	fmt.Println(a, b)

	prg2()
	prg3()
}

func prg2(){
	c :=2
	fmt.Printf("%d\t%b\t%#x\n",c,c,c,c,)

	d := c<<1
	fmt.Println("This is shifting bits to right \t",d)

}
const(
	e = 2005
	f = 2005 + iota
	g = 2005 + iota
)

func prg3(){
	fmt.Println(e,f,g)
}