package main

import (
	"fmt"
	// "slices"
	// "slices"
)

func sliceFunc() {
	// [] is empty for slices [...] makes it an array
	// var sl1 = [] int {10, 20, 30};
	// var sl2 = [] int {1: 5, 5: 10, 6}
	// var sl3 [] int;
	// fmt.Println("sl1: ", sl1);
	// fmt.Println("sl2: ", sl2);
	// fmt.Println("sl3: ", sl3);
	// fmt.Println(sl3 == nil);
	// fmt.Println(sl2 == nil);
	// fmt.Println(slices.Equal(sl1, sl2));
	// sl3 = append(sl3, 10);
	// fmt.Println(sl3, len(sl3), cap(sl3));
	// sl2 = append(sl2, 11);
	// fmt.Println(sl3, len(sl3), cap(sl3));
	// sl3 = append(sl3, 13, 14, 15);
	// fmt.Println(sl3, len(sl3), cap(sl3));
	// sl3 = append(sl3, 13, 14, 15);
	// fmt.Println(sl3, len(sl3), cap(sl3));
	// sl3 = append(sl3, 13, 14, 15);
	// fmt.Println(sl3, len(sl3), cap(sl3));

	// var sl4 = make([]int, 5, 5) // make(type, len, cap)
	// sl4 = append(sl4, 10, 20, 30, 40, 50)
	// fmt.Println(len(sl4), sl4, cap(sl4))
	// sl4 = append(sl4, 60)
	// fmt.Println(len(sl4), sl4, cap(sl4))

	//emptying a slice
	// var sl5 = []string{"apple", "ball", "cat", "dog", "egg"}
	// fmt.Println("sl5: ", sl5, len(sl5))
	// clear(sl5)
	// fmt.Println("sl5: ", sl5, len(sl5))

    // zero length slice, it doesn't equal to nil
    // var sl6 = [] int {};
    // var sl7 [] int;
    // var sl8 [] int;
    // fmt.Println("zero length slice: ", sl6);
    // fmt.Println(slices.Equal(sl6, sl7));
    // fmt.Println(slices.Equal(sl8, sl7));
    // fmt.Println(sl6 == nil); // false
    // fmt.Println(sl7 == nil); // true

    // slicing slices
    // var sl9 = []string {"a", "b", "c", "d"};
    // var sl10 = sl9[:2]
    // var sl11 = sl9[1:3]
    // var sl12 = sl9[:]
    // sl9[0] = "A"
    // sl12[1] = "B"
    //
    // fmt.Println(sl10)
    // fmt.Println(sl11)
    // fmt.Println(sl12)

    // copy
    // var sl13 = [] int {1, 2, 3}
    // var sl14 = make ([] int, 5)
    // num := copy(sl14, sl13)
    // fmt.Println(num, sl14)

    var sl15 = [] int {10, 20, 30, 40, 50}
    var sl16 = make([]int, 3)
    num := copy(sl16, sl15[2:])
    sl15[3] = 101
    // num = copy(sl16[:2], sl15[1:])
    fmt.Println(num, sl16)


}
