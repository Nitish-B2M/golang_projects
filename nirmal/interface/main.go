package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	intrface := []interface{}{1, 3.23, "nitish"}
	fmt.Println(intrface)

	nirmal := Student{Name: "nirmal", Age: 25}
	nitish := Student{Name: "nitish", Age: 23}
	fmt.Println(nirmal, "Age:", nirmal.Age)
	fmt.Println(nitish, "Age:", nitish.Age)
}
