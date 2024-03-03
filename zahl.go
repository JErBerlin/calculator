package main

import (
	"fmt"
	"strconv"
)

// Num ist ein eskalarer Wert: int in dem Fall
type Num int

// Zahl ist der (numerischer) Ausdruck einer Zahl
type Zahl struct {
	n Num
}

// Wert gibt den numersichen Wert der Zahl zurück
func (z Zahl) Wert() (Num, error) {
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
func (z *StrZahl) Wert() (Num, error) {
	if z.zahlAlsString == "" {
		return 0, fmt.Errorf("Wert von StrZahl: Zeichenkette ist leer")
	}
	i, err := strconv.Atoi(z.zahlAlsString)
	return Num(i), err
}

func (z *StrZahl) Reset() {
	z.zahlAlsString = ""
}
