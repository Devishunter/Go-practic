package main

import "fmt"

func main() {
	/*integers := make([]int, 10)
	fmt.Println(integers)

	for i := range integers {
		integers[i] = i

	}

	fmt.Println(integers)*/

	sammy := "Sammy"

	for _, letter := range sammy {
		fmt.Printf("%c\n", letter)
	}
}
