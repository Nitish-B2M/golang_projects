package main

import (
	"fmt"
)

var ji int = 2

// go statically typed language
func main() {
	// 1.  normal variable declaration
	// var abc int16 = 16561
	// var str string = "Hii"
	// // 2. short hand declaration
	// def := 1
	// // 3. multiple variable declaration
	// a, b, c := 1, true, "character"
	// // const
	// const MAX = math.MaxInt

	// char, rune, untyped int

	// fmt.Println(abc, def, a, b, c, MAX, ji)
	// fmt.Printf("%v %d %T\n", str, a, a) // %v: print any data type values, %T: used to print actual data type

	// Map {"nirmal":true}, {"apple": 4}
	// 1. normal declare
	maps := map[string]bool{"nirmal": true, "nitish": false}

	// 2.
	maps2 := make(map[string]int) // data-type, length, cap
	maps2["apple"] = 4
	// fmt.Println(maps, maps2)
	fmt.Printf("%+v\n", maps)

	find_value, ok := maps["nirmal"]
	if !ok {
		fmt.Println("Key not exist")
	} else {
		fmt.Println(find_value)
	}

	// Array(fixed data-type)/Slice(dynamic data-type)
	var arr [2]int
	arr[0] = 12
	arr[1] = 34
	fmt.Println(arr)

	arr2 := [3]string{"nirm", "nit", "sou"}
	fmt.Println(arr2)

	slce1 := []string{"nirm", "nit", "sou"}
	// fmt.Printf("%T %T", arr2, slce1)
	fmt.Println(slce1, len(slce1), cap(slce1))

	slce2 := make([]int, 1, 4)
	fmt.Println(slce2, len(slce2), cap(slce2))
	slce2[0] = 19             // 1st index
	slce2 = append(slce2, 10) // on top of length 2 appending 10 // index 2nd
	slce2 = append(slce2, 11)
	slce2 = append(slce2, 3)
	slce2 = append(slce2, 9)
	fmt.Println(slce2, len(slce2), cap(slce2))
	slce2 = append(slce2, 911)
	slce2 = append(slce2, 311)
	slce2 = append(slce2, 922)
	slce2 = append(slce2, 921)
	fmt.Println(slce2, len(slce2), cap(slce2))

	// function
	// Method
}
