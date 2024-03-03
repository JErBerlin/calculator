package main

import (
	"io"
)

// Ausdr repräsentiert einen Ausdruck, der ausgewertet werden kann.
type Ausdr interface {
	Wert() (Num, error)
}

// baueAusdr takes an io.Reader, parses the input into an Ausdr, and returns it.
func baueAusdr(r io.Reader) (Ausdr, error) {
	scanner := NewSymbolScanner(r)
	return ParseAusdruck(scanner)
}
