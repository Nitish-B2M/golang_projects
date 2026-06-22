package anonymous

import "fmt"

func Anonymous() {
	// anonymous
	xyz := func() int {
		fmt.Println("Do nothing")
		return 123
	}

	cer := xyz()
	fmt.Println("cer", cer)
}

func common() int {
	abc := 3
	return abc
}

func named() (abc int) {
	abc = 3
	return
}
