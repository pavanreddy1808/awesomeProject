package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age   int
}
	type jsperson struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Age   int    `json:"Age"`
}


func main(){

	p1:= person{
	Fname: "Pavan Reddy ",
	Lname: "Kumar",
	Age: 29,
	}
	p2:=person{
		Fname: "Navya",
		Lname:"Pabbathi",
		Age: 25,
	}

	user:=[]person{p1,p2}        // variable is of type [] struct using p1 & p2
	fmt.Println(user)
	bs,err:=json.Marshal(user)    // Marshaling
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))     //bs is of type []byte

	unjsperson:= []jsperson{}    // var unjsperson []jsperson is another way of creating

	err =json.Unmarshal(bs, &unjsperson)   //unmarshal accepts []byte
	if err!= nil{

		fmt.Println(err)
	}
	fmt.Println(unjsperson)
}