package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
}

/*page 27*/
/* Донован, Алан А. А., Керниган, Брайан, У.
Д67 Язык программирования Go. : Пер. с англ. — М. : ООО “И.Д. Вильямс”,
2016. — 432 с. : ил. — Парал. тит. англ.*/
