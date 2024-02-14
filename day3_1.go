package main

import "fmt"

func main() {
	sharks := []string{"hammerhead", "dogfish"}

	for i := 0; i < len(sharks); i++ {
		fmt.Println(sharks[i])
	}

	for i, sharks := range sharks {
		fmt.Println(i, sharks)
	}

	for _, sharks := range sharks {
		fmt.Println(sharks)
	}

	for range sharks {
		sharks = append(sharks, "shark")
	}

	fmt.Println("\n", sharks)
}
