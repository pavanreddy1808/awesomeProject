package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

	//encoding go objects
	err = json.NewEncoder(os.Stdout).Encode(user)
	if err!=nil{
		fmt.Println(err)
	}
	//fmt.Println(user)

	//sorting
	s:=[]string{"pavan","reddy","Navya","pabbathi"}
	i:=[]int{2,4,7,2,4,6,8,1}
	fmt.Println(s)
	fmt.Println(i)
	//after sort
	sort.Strings(s)
	sort.Ints(i)
	fmt.Println(i)
	fmt.Println(s)
}