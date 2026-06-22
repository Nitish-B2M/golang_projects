package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	// TODO: Implement the function
	// in s we have string with x length,
	// so first get the length
	// then use 2 pointer to reverse it

	rStr := []rune(s)
	fmt.Println(rStr)
	j := len(rStr) - 1
	for i := 0; i < j; i++ {
		temp := rStr[j]
		rStr[j] = rStr[i]
		rStr[i] = temp
		j--
	}

	return string(rStr)
}
