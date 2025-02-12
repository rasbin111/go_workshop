/*
	Type constraints allow you to specify the list of data types that you want to work with

in order to avoid logical errors and bugs
*/
package main

import "fmt"

func Same[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func constraintsFunc() {
	fmt.Println("4 = 3 is", Same(4, 3))
	fmt.Println("aa = aa is", Same("aa", "aa"))
	fmt.Println("4.1 = 4.15 is", Same(4.1, 4.15))
    fmt.Println("Comparing two arrays: ", Same([...]int {10, 20, 30}, [...]int {10, 20, 30}))
}
