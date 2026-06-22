package main

import "fmt"

func normalStruct() {
	//  normal struct
	type Person struct {
		Name string
		Age  int
	}

	nirmal := Person{Name: "nirmal", Age: 25}
	fmt.Println(nirmal, "Age:", nirmal.Age)
	nirmal.Age = 23
	fmt.Println(nirmal, "Age:", nirmal.Age)
}

func nestedStruct() {
	//  normal struct
	type School struct {
		sName string
	}

	type Person struct {
		Name string
		Age  int
		School
	}

	nirmal := Person{Name: "nirmal", Age: 25, School: School{sName: "Ganpat"}}
	fmt.Println(nirmal, "Age:", nirmal.Age)
	nirmal.Age = 23
	nirmal.School.sName = "GNUI"
	fmt.Println(nirmal, "Age:", nirmal.Age)
}

func main() {
	normalStruct()
	nestedStruct()
}
