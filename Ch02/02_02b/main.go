package main

import (
	"fmt"
)

func main() {

	 str1 := "The quick red fox"
	 str2 := "jumped over"
	 str3 := "the lazy brown dog."
	 aNumber := 47

	fmt.Println(str1, str2, str3)
	strLenght, err := fmt.Println("The value is: ", aNumber)
	if err == nil {
		fmt.Println("String length:", strLenght)
	}
	fmt.Printf("Value of the number: %v\n", aNumber)
	fmt.Printf("Data type: %T\n", aNumber)
}
