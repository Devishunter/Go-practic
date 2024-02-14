package main

import "fmt"

func main() {
	var favoritColor = "red"
	fmt.Println("Color is", favoritColor)

	birthYear, ageInYears := 1998+, 25
	fmt.Println("Родился в", birthYear, "году", "мне", ageInYears, "года")

	var (
		firstInitial   = "A"
		surnameInotial = "B"
	)

	fmt.Println("Инициалы =", firstInitial, surnameInotial)

	var ageInDays int
	ageInDays = 365 * ageInYears

	fmt.Println("Мой возраст в днях =", ageInDays)
}
