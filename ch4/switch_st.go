package main

import "fmt"

func switchFunc(){
    /*
    words := [] string {"a", "cow", "smile", "laugh", "octopus", "whale", "occupation"}
    for _, word := range words{
        switch size := len(word); size {
        case 1, 2, 3, 4:
            fmt.Println(word, "is a short word")
        case 5:
            fmt.Println(word, "is a medium word")
        case 6, 7, 8, 9:
        default:
            fmt.Println(word, "is a long word")
        }
    } 
    */

    // blank switch:
    // a regular switch: allows you to check a value for equality
    // a blank switch: allows you to use any boolean comparison 
    words := [] string {"a", "cow", "smile", "laugh", "octopus", "whale", "occupation"}

    for _, word := range words {
        switch size := len(words); {
        case size < 5 && size > 2:
            fmt.Println(word, "is a short word")
        case size > 8:
            fmt.Println(word, "is a long word") 
        default:
            fmt.Println(word, "is a perfect word size")

        }
    }

}
