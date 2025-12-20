package cycles

import "fmt"

func Cycles() {

	str := "Hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

}

func sumNegativeNumbers() {
	numbers := [...]int{-1, -2, -3, -4, -5, 6, 7, 8, 9, 10}
	sum := 0
	for _, v := range numbers {
		if v < 0 {
			fmt.Println("Negative number found:", v)
			continue
		}
		sum += v
	}
	fmt.Println("Sum of non-negative numbers:", sum)
}
