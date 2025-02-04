package main

import (
	"fmt"
	"math/rand"
)

func ifFunc(){
    if age := rand.Intn(30); age <18{
        fmt.Println("You can't vote, because you are", age)
    } else if age == 18{
        fmt.Println("Check your status")
    } else {
        fmt.Println("You are", age, "and can vote")
    }

}
