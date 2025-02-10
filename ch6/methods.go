package main 

import "fmt"

type Citizen struct {
    nationality string
    fname string
    lname string
    age int8
}

// method for Citizen struct 
func (c Citizen) fullName() string {
    return c.fname + " " + c.lname
}


func methodFunc(){
    c1 := Citizen{
        nationality: "nepali",
        fname: "ram",
        lname: "thapa",
        age: 10,
    }
    fullName := c1.fullName()
    fmt.Println(fullName)
    
}
