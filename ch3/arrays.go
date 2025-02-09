package main

import (
	"fmt"
)

func arrayFunc() {
	var arr1 [3]int
	var arr2 []int
	var arr3 [3]int = [3]int{10, 20, 30}
	var arr4 = [3]int{50, 60, 70}
	var arr5 = [10]int{1: 11, 5: 55}
	var slice6 = []int{10, 20, 30} // its a slice not array
	var arr7 = [...]int{50, 60, 07, 65}
	// var arr8  = {10, 20, 30, 50}; // error

	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr3: ", arr3)
	fmt.Println("arr4: ", arr4)
	fmt.Println("arr5: ", arr5)
	fmt.Println("slice6: ", slice6)
	fmt.Println("arr7: ", arr7)
	fmt.Println("Len arr5: ", len(arr5))

	// var compareArray = arr4 == arr5; // error as array of differnt size is considered different type in go
	// var compareArray = arr4 == arr6; // this also differnt type as arr4 is [3] int and arr6 is [] int
	var compareArray = arr1 == arr3
	fmt.Println("array comparison: ", compareArray)

}
