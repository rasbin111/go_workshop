package main 

import "fmt"

func deferFunc() {
    a := 10
    defer func (val int) {
        fmt.Println("first:", val)
    }(a)

    a = 20
    defer func (val int) {
        fmt.Println("second:", val)
    }(a)

    a = 30
    fmt.Println("Existing", a)

}
