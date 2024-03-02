package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// Ausdr repräsentiert einen Ausdruck, der ausgewertet werden kann.
type Ausdr interface {
	Wert() (int, error)
}

// baueAusdr baut einen Ausdruck aus einer Rune-Zeichenkette im Reader
func baueAusdr(r io.Reader) (Ausdr, error) {
	reader := bufio.NewReader(r)

	var exp Ausdr
	var op Verkn
	var str StrZahl // String um die gelesenen Ziffer zu speichern und dann zum Nummerischen Wert zu konvertieren

	// lese eine grupe der Form A | A + B | A - B | A * B | A / B
	// noch keine Parenthesis erlaubt
	for {
		symbol, _, err := reader.ReadRune()

		// EOF error ist kein eigentlicher Fehler
		if err == io.EOF {
			break // Dateiende
		}
		// sonstiger Fehler, stop
		if err != nil {
			return exp, err
		}
		// sortiere Leerzeichen aus
		if unicode.IsSpace(symbol) {
			continue
		}

		if unicode.IsDigit(symbol) { // lese eine Ziffer von Operanden a oder Operanden b nin die Stringzahl
			str.Hinzu(symbol)
			continue
		}

		// kein leerzeichen und keine Ziffer: versuche, eine Verknupfung zu lesen
		op, err = leseVerknupfung(symbol)
		if err != nil {
			return exp, fmt.Errorf("leseVerkn: fehler beim Lesen der Verknüpfung: %s", err)
		}

		// speichere erstem Operanden A in die Verknupfung
		w, err := str.Wert()
		if err != nil {
			return exp, fmt.Errorf("baueAusdr: vor der Verknüpfung erwarteter Operanden A nicht vorhanden")
		}
		op.SetA(w)

		// speichere die Verknupfung als Ausdrueck exp
		exp = op

		// reset Zahlenleser, um Operanden b zu lesen
		str.Reset()
	}

	// Die Zeichenkette eintält eine Zahl alleine oder den zweiten Operanden B
	w, err := str.Wert()
	if err != nil { // etwas ist schiefgelaufen: wir brauchen wenigstens eine Zahl
		return exp, fmt.Errorf("baueAusdr: die Zeichenkette enthält keine Zahl")
	}

	// wenn der Ausdrueck eine Verknuepfung ist, speichern wir B
	if v, ok := exp.(Verkn); ok {
		v.SetB(w)
	} else { // sonst speichern wir eine Zahl alleine
		exp = Ganzezahl{w}
	}

	return exp, nil
}
