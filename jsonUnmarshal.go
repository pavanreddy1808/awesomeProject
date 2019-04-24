package main

import (
	"encoding/json"
	"fmt"
)
type p3 struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Age   int    `json:"Age"`
}
type person struct {
	Fname string
	Lname string
	Age int
}
func main() {

	p1:= person{
		Fname:"Pavan Reddy ",
		Lname:"reddy",
		Age: 32,
	}
	p2:=person{
		Fname: "Navya Reddy",
		Lname: "Pabbathi",
		Age: 25,
	}

people:=[]person{p1,p2}  // Marshal function returns []bytes hence we created one
fmt.Println(people)
bs, err :=json.Marshal(people)   //marshalling string into json object    go object into json object
if err!=nil {
	fmt.Println(err)  // printing error
}
	fmt.Println(string(bs))       //Printing json object by converting into string
	js := fmt.Sprint(string(bs))   //string print and assinging value to variable so that we can use it to unmarshal
	fmt.Println("ja value",js)
ums:=[]byte(js)
var unm []p3     //unm:=[]p3{} another way
fmt.Println("ums value",ums)
fmt.Printf("\n%T\t",ums)
err =     json.Unmarshal(ums, &unm)  //unmarshal takes only []byte    and unmarshal into our structure of kind p3  taking jason object into go object
if err!=nil{

	fmt.Println(err)
}
fmt.Println("unm value",unm)
fmt.Printf("%T",unm)
	fmt.Printf("%T",ums)
	fmt.Printf("%T",js)
	fmt.Printf("%T",p1)
}

