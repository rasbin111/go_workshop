package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
    Name string
    Body string 
    Time int64 
}

func main (){

    m1 := Message{"Alice", "Hello", 8282918}

    b, err := json.Marshal(m1)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%s \n", b)

    var m2 Message 
    jerr := json.Unmarshal(b, &m2)
    if jerr != nil {
        fmt.Println(err)
    }
    fmt.Println(m2)
    fmt.Printf("Type of m2: %T \n", m2)


}
