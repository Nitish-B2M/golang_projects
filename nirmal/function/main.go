package main

import (
	"fmt"
	"function/anonymous"
)

// normal function
func sum() {
	// Addition
	a, b := 1, 2
	sum := a + b
	fmt.Println("Sum:", sum)
}

// parameterized function
func sub(a, b int) {
	// Addition
	sub := a - b
	fmt.Println("Sub:", sub)
}

// return  function
func mul(a, b int) int {
	// Addition
	mul := a * b
	return mul
}

// return function with multiple parameter
func mulDiv(a, b int) (int, int) {
	// Addition
	mul := a * b
	div := b / a
	return mul, div
}

// pass by value
func passByValue(a int) {
	fmt.Println("[Func] pass by value:", a)
	a = 98
	fmt.Println("[Func] pass by value:", a)
}

// pass by reference
func passByReference(b *int) {
	fmt.Println("[Func] pass by reference:", *b)
	*b = 98
	fmt.Println("[Func] pass by reference:", *b)
}

func isEven(val int) (bool, error) {
	if val%2 != 0 {
		return false, fmt.Errorf("this is not an even number")
	}
	return true, nil
}

func main() {
	// sum()
	// sub(4, 3)

	// mul
	// val := mul(3, 8)
	// fmt.Println("Mul:", val)

	// mul div
	// val1, val2 := mulDiv(3, 8)
	// fmt.Println("Mul:", val1)
	// fmt.Println("Div:", val2)

	// // pass by value
	// a := 2
	// passByValue(a)
	// fmt.Println("[Main] pass by value:", a)

	// // pass by reference
	// b := 10
	// passByReference(&b)
	// fmt.Println("[Main] pass by reference:", b)

	// if ok, err := isEven(3); ok {
	// 	fmt.Println("it's an even number")
	// } else {
	// 	fmt.Println(err)
	// }

	// // calling add function from calculator module
	// calculator.Add()

	anonymous.Anonymous()
}
