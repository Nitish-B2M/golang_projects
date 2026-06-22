package calculator

import (
	"fmt"
	"function/anonymous"
)

func Add() {
	a, b := 1, 2
	sum := a + b
	fmt.Println("[Calculator] Sum:", sum)

	anonymous.Anonymous()
}
