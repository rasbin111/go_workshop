package main

import (
	"fmt"
	"sort"
)

type Student struct {
	firstName string
	lastName  string
	age       int8
}

// returning function from function
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func closureFunc() {

	people := []Student{
		{"rgt", "thapa", 24},
		{"sbgt", "thapa", 25},
		{"rbgt", "thapa", 27},
	}

	// sort by firstName
	sort.Slice(people, func(i, j int) bool {
		return people[i].firstName < people[j].firstName
	})
	fmt.Println(people)
	baseTwo := makeMult(2)
	baseThree := makeMult(3)

	fmt.Println(baseTwo(9))
	fmt.Println(baseThree(9))

}
