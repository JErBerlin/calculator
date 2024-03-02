package main

import (
	"fmt"
	"strings"
)

func main() {
	strAusdr := "1 + 2"
	rd := strings.NewReader(strAusdr)

	a, err := baueAusdr(rd)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	ergebnis, err := a.Wert()
	if v, ok := a.(Verkn); ok {
		fmt.Printf("Ausdruck ist die Verknuepfung: %d %s %d\n", v.A(), string(v.Operation()), v.B())
		if err != nil {
			fmt.Printf("Wert(%s) = error: %s", strAusdr, err)
		} else {
			fmt.Printf("Wert(%s) = %d\n", strAusdr, ergebnis)
		}
	} else if _, ok := a.(Ganzezahl); ok {
		fmt.Println("Ausdruck ist eine Ganzezahl")
		if err != nil {
			fmt.Printf("Wert(%s) = error: %s", strAusdr, err)
		} else {
			fmt.Printf("Wert(%s) = %d\n", strAusdr, ergebnis)
		}
	} else {
		fmt.Println("Ausdruck war unlesbar. Unfug.")
	}
}
