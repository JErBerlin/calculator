package main

import (
	"fmt"
	"strconv"
)

// Ganzezahl ist der (numerischer) Ausdruck einer Ganzezahl
type Ganzezahl struct {
	n int
}

// Wert gibt den numersichen Wert der Ganzezahl zurück
func (z Ganzezahl) Wert() (int, error) {
	return z.n, nil
}

// StrZahl ist ein Typ, der eine Zeichenkette zur Zahlendarstellung verwendet.
type StrZahl struct {
	zahlAlsString string // Zeichenkettendarstellung der Zahl
}

// Hinzu fügt eine Ziffer zur Zeichenkette hinzu.
func (z *StrZahl) Hinzu(d rune) {
	z.zahlAlsString += string(d)
}

// Wert konvertiert die Zeichenkettendarstellung der Zahl in eine Ganzzahl.
func (z *StrZahl) Wert() (int, error) {
	if z.zahlAlsString == "" {
		return 0, fmt.Errorf("Wert von StrZahl: Zeichenkette ist leer")
	}
	i, err := strconv.Atoi(z.zahlAlsString)
	return i, err
}

func (z *StrZahl) Reset() {
	z.zahlAlsString = ""
}
