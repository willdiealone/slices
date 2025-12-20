package uslOperatrori

import "fmt"

func IfElseSwitch() {

	a := 2
	b := 2

	if a < b {
		fmt.Println("a is less than b")
	} else if a > b {
		fmt.Println("a is not less than b")
	} else {
		fmt.Println("a is equal to b")
	}

	switch a < b {
	case true:
		fmt.Println("a is less than b")
		fallthrough
	case false:
		fmt.Println("fallthrough")
		fmt.Println("a is equal to b")
	default:
		fmt.Println("a is not less than b")
	}
}
