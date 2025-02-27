package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
    Name string `json:"username"`
    Age int32 `json:"age"`
    Phone string `json:"phone_number"`
    Address string `json:"address"`
}

func check (e error){
    if e != nil {
        log.Fatal(e)
    }
}
func ReadJsonFile(){
    data, err := os.ReadFile("./names.json")
    check(err)
    var p Person

    err = json.Unmarshal(data, &p)
    check(err)
    fmt.Println(p)
    
}
