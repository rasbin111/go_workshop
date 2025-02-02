package main

import "fmt"

func CMain(){
    const a = 10;
    const b string = "apple";
    const c rune = 'b';
    const d float32 = 10.567;
    const e = true;
    fmt.Printf("a: %d\n", a);
    fmt.Printf("b: %s\n", b);
    fmt.Printf("c: %c\n", c);
    fmt.Printf("d: %.2f\n", d);
    fmt.Printf("e: %t\n", e);
}
