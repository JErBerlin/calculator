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

	var aktuelleVerkn Verkn
	var strZahl StrZahl // String um die gelesenen Ziffer zu speichern und dann zum Nummerischen Wert zu konvertieren

	// lese eine grupe der Form A | A + B | A - B | A * B | A / B
	// noch keine Parenthesis erlaubt
	for {
		symbol, _, err := reader.ReadRune()

		if err == io.EOF {
			break // Dateiende
		}
		// sonstiger Fehler, stop
		if err != nil {
			return nil, fmt.Errorf("Fehler beim Lesen: %v", err)
		}

		// sortiere Leerzeichen aus
		if unicode.IsSpace(symbol) {
			continue
		}

		// lese eine Ziffer von Operanden a oder Operanden b in die Stringzahl
		if unicode.IsDigit(symbol) {
			strZahl.Hinzu(symbol)
			continue
		}

		// kein leerzeichen und keine Ziffer: versuche, eine Verknupfung zu lesen
		if aktuelleVerkn, err = leseVerknupfung(symbol); err != nil {
			return nil, fmt.Errorf("leseVerkn: Fehler beim Lesen der Verknüpfung: %s", err)
		}

		// speichere Operanden A in die Verknupfung
		zahl, err := strZahl.Wert()
		if err != nil {
			return nil, fmt.Errorf("baueAusdr: Fehler beim Konvertieren Zeichen zu Zahl für Operanden")
		}
		aktuelleVerkn.SetA(zahl)

		// bereit für Operanden B
		strZahl.Reset()
	}

	// Die Zeichenkette eintält eine Zahl alleine oder den zweiten Operanden B
	letzteZahl, err := strZahl.Wert()
	if err != nil {
		// etwas ist schiefgelaufen: wir brauchen wenigstens eine Zahl
		return nil, fmt.Errorf("die Zeichenkette enthält keine Zahl: %s", err)
	}

	// der Ausdrueck ist entweder eine Zahl oder eine Verknupfung
	if aktuelleVerkn == nil {
		return Ganzezahl{letzteZahl}, nil
	}

	aktuelleVerkn.SetB(letzteZahl)
	return aktuelleVerkn, nil
}
