package main 

import "fmt"

type Employee struct {
    age int 
    name string 
}

func modifyFails(i int, s string, e Employee){
    i = i * 2
    s = "Goodbye"
    e.name = "Unknown"
}

func modMap(m map[int]string){
    m[2] = "Hello"
    m[3] = "Goodbye"

    delete(m, 1)
}

func modSlice (s []int){
    for k, v := range s{
        s[k] = v * 2
    }

    s = append(s, 10)
}

func CbvFunc(){
    e := Employee{
        10,
        "rajan",
    }
    i := 10
    s := "Hello"

    modifyFails(i, s, e)
    fmt.Println(i, s, e)

    sl1 = [] int {10, 20, 30, 40, 50}
    mp1 = {
        1: "apple",
        2: "ball",
        3: "cat"
    }

    modMap(mp1)
    fmt.Println(mp1)

    modSlice(sl1)
    fmt.Println(sl1)
}
