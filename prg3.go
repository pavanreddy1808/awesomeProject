package main
import "fmt"
// Using TYPE in Go
// creating our own TYPE in GO
var a int
type chicken int
func main(){
a := 42
var b chicken = 30
fmt.Println(a)
fmt.Println(b)
a = int(b)
fmt.Println(a)
}