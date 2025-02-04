package main

import (
	"fmt"
	"math/rand"
)

func gotoFunc(){
    a := rand.Intn(10)
    for a < 100 {
        if a % 5 == 0{
            goto done
        }
        a = a * 2 + 1
    }
    fmt.Println("loop completed normally")
done:
    fmt.Println("loop ended somewhere before end")
    fmt.Println(a)
}
