package main

import (
	"fmt"
	"strings"
)

func main() {
	strAusdr := "1 + 2"
	rd := strings.NewReader(strAusdr)

	ausdruck, err := baueAusdr(rd)
	if err != nil {
		fmt.Printf("Fehler: %v\n", err)
		return
	}

	wert, err := ausdruck.Wert()
	if err != nil {
		fmt.Printf("Fehler beim Auswerten: %v\n", err)
		return
	}

	fmt.Printf("Wert(%s) = %d\n", strAusdr, wert)
}
