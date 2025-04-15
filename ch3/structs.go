package main

import "fmt"

func StructDemo() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	fmt.Println("Fred Person", fred)

	bob := person{}
	bob.name = "bob"
	fmt.Println(bob)

	julia := person{
		"julia",
		18,
		"dog",
	}

	fmt.Println(julia)

	beth := person{
		name: "beth",
		age:  20,
	}
	fmt.Println(beth)

}
