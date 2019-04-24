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

func main() {

	p1 := person{
		Fname: "Pavan Reddy ",
		Lname: "reddy ",
		Age:   29,
	}

	p2 := person{
		Fname: "Navya",
		Lname: "reddy",
		Age:   25,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
