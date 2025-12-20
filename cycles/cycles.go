package cycles

import "fmt"

func Cycles() {

	str := "Hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

}

func SumNegativeNumbers() {
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

func SumNegativeNumbersBreak() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, -8, 9, 10}
	sum := 0
	for _, v := range numbers {
		if v < 0 {
			fmt.Println("Negative number found:", v)
			break
		}
		sum += v
	}
	fmt.Println("Sum of non-negative numbers before negative:", sum)
}

func MultiplicationTable() {
	for i := 1; i < 10; i++ {
		if i == 9 {
			break
		}
		for j := 1; j < 10; j++ {

			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

}
