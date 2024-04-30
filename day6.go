package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}

/* Донован, Алан А. А., Керниган, Брайан, У.
Д67 Язык программирования Go. : Пер. с англ. — М. : ООО “И.Д. Вильямс”,
2016. — 432 с. : ил. — Парал. тит. англ.*/
