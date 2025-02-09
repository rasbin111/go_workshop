package main

import (
	"errors"
	// "fmt"
	// "os"
)

func div(num int, denom int) int { // can do this for param (num, denom int)
    if denom == 0 {
        return 0
    }
    return num / denom
}

// does not allow named and optional parameters, so using struct to add named and optional parameters
type Person struct{
    fname string
    mname string
    lname string 
    age int
    address string
}

func returnFname(p1 Person) string {
    return p1.fname + " " + p1.mname + " " + p1.lname 
}

// variadic input parameters
func addTo(base int, vals ...int) []int {
    out := make([]int, 0, len(vals))

    for _, v := range vals{
        out = append(out, base+v)
    }
    return out
}

// variadic input parameters
func sum(vals ...int) int {
    sum := 0
    for _, v := range vals{
        sum += v
    }
    return sum
}

// mutiple return values
func divAndRemainder(num, denom int) (int, int, error){
    if denom == 0{
        return 0, 0, errors.New("can't divide by zero")
    }
    return num /denom, num % denom, nil
}

func functionFunc(){
    // var div1 int = div(40, 20)
    // fmt.Println(div1)
    // fmt.Println(addTo(3, 10, 20, 30, 40))
    // fmt.Println(sum(10, 20, 20, 50))
    // fmt.Println(returnFname(Person{
        // fname: "ram",
        // mname: "bdr",
        // lname: "thapa",
        // age: 10,
    // }))
    // div, rem, err := divAndRemainder(30, 7)
    // div2, rem2, err2 := divAndRemainder(30, 0)
    // if err2 != nil {
    //     fmt.Println(err2)
    //     os.Exit(1)
    // }
    // fmt.Println(div2, rem2)

    divAndRemainder(20, 10)
}
