package main 

import "fmt"

func PrintSlice[T any] (s []T) {

    for _, v := range s {
        fmt.Println(v)
    }
}
func genericsFunc(){
    sl1 := [] int {10, 20, 30, 40, 50}
    PrintSlice(sl1)
    sl2 := [] string {"apple", "ball", "cat"}
    PrintSlice(sl2)

}
