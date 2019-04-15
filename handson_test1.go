package main

import "fmt"
var x int
var y string
var z bool
type mutton int
func main(){
 a := 42
 b := "James Bond"
 c := true
 fmt.Println("hands on 1")
 fmt.Println(a, b, c)
 fmt.Printf("%T\n%S\n%B\n",a, b, c)
	fmt.Println("hands on 2")
 hnd2()
	fmt.Println("hands on 3")
	hnd3()
	fmt.Println("hands on 4")
	hnd4()
	fmt.Println("hands on 5")
	hnd5()
}
// Using Global Variables (package level)
func hnd2(){
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
// assigning value using Sprint
func hnd3(){
	x = 007
	y = "James Bond"
	z = false
	fmt.Printf("%v \t%v\t%v\t",x,y,z)
	s := fmt.Sprintf("%v \t%v\t%v\t",x,y,z)
	fmt.Println(s)
}
//Using TYPE
func hnd4(){
	type chicken int
	var a chicken
	fmt.Println(a)
	fmt.Printf("%T \t", a)
	a = 20
	fmt.Println(a)
}
//conversions
func hnd5(){

	var a mutton
	var c int
	c = 10
	fmt.Println(c)
	fmt.Println(a)

	a = mutton(c)
	fmt.Println(a)

}