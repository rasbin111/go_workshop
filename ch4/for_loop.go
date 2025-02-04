package main

import (
	"fmt"
)

func forFunc(){
    // var mySlice = []int{1, 2, 3, 4, 5}
    // mySliceTwo := [] int {4, 5, 6, 7, 9}
    
    // for-range
    // for i, v := range mySlice{
    //     fmt.Println(i, v)
    // }

    // for-range
    // for i, v := range mySliceTwo{ // we can ommit second value
    //     fmt.Println(i, v)
    // }

    // for loop
    // for i := 1; i<=10; i++{
    //     fmt.Println(i)
    // }

    // condition only
    // i := 2
    // for i <= 10 {
    //     fmt.Println(i);
    //     i += 2
    // }

    // infinite for - loop 
    // i := 2
    // for {
    //     if i > 10{
    //         break
    //     }
    //     fmt.Println(i)
    //     i++
    // }

    // for-range with maps
    // uniqueNames := map[string]bool {"rgt": true, "ram": false, "shyam": false, "god": true}
    //
    // for k, v := range uniqueNames { // we can omit v as we can only use k 
    //     fmt.Println(k, v)
    // }

    // for-range value is a copy
    // scores := [] int {88, 87, 79, 97, 77}
    // for _, v := range scores {
    //     v *= 2
    // }

    // nested for with label 
    words := [] string {"hello", "apple"}
outer:
    for _, v := range words {
        for i, iv := range v {
            fmt.Println(i, iv, string(iv))
            if iv == 'l'{
                continue outer
                // break
            }
        }
    }



    
}
