package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// Zeichenkette, die ausgewertet werden soll
	strAusdr := "3*(1+2)"

	r := strings.NewReader(strAusdr)
	br := bufio.NewReader(r)

	// parse Ausdruck
	ausdr, err := ParseAusdruck(br)
	if err != nil {
		fmt.Printf("Fehler beim Parsen des Ausdrucks: %s\n", err)
		return
	}

	// Auswertung des Ausdrucks
	w, err := ausdr.Wert()
	if err != nil {
		fmt.Printf("Fehler bei der Auswertung des Ausdrucks: %s\n", err)
		return
	}

	fmt.Printf("Wert(%s) = %d\n", strAusdr, w)
}
