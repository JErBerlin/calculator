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

// ParseAusdruck baut einen Ausdruck rekursiv aus einer Zeichenkeete r
func ParseAusdruck(r *bufio.Reader) (Ausdr, error) {
	a, err := ParseTeilAusdruck(r, 0)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// ParseTeilausdruck vearbeitet Teil-Ausdruecke basierend auf der Prioritaet der Operationen
func ParseTeilAusdruck(r *bufio.Reader, minPrio int) (Ausdr, error) {

	// Als erstes Parse linke Ausdruck, bis zum Klammer: z.B. (links) + (rechts)
	links, err := ParseGrundAusdruck(r)
	if err != nil {
		return nil, err
	}

	// rekusive Schleife
	// von links nach rechts und durch prio
	for {
		op, err := leseRune(r)
		if err != nil {
			if err == io.EOF {
				return links, nil
			}
			return nil, err
		}

		prio := OpPrio(op)
		if prio < minPrio { // op gehoert einer höheren Ebene der Rekursion / des Baums
			r.UnreadRune()    // -> leg das Symbol zurück
			return links, nil // Ausdruck besteht nur aus dem linken Teil
		}

		rechts, err := ParseTeilAusdruck(r, prio+1) // rekursives Aufrufen im rechten Ausdruck
		if err == io.EOF {
			return links, nil
		}
		if err != nil {
			return nil, err
		}

		verkn, err := leseVerknupfung(op)
		if err != nil {
			return nil, err
		}

		// Auswertung von links-Ausdr A und rechts-Ausdr B
		linksWert, err := links.Wert()
		if err != nil {
			return nil, err
		}
		verkn.SetA(linksWert)

		rechtsWert, err := rechts.Wert()
		if err == io.EOF {
			return links, nil
		}
		if err != nil {
			return nil, err
		}
		verkn.SetB(rechtsWert)

		links = verkn // Setze aktuelle Verknüpfung als neuen linken Operanden
	}

}

// ParseGrundAusdruck liest und verarbeitet Zahlen und geklammerte Ausdruecke.
// Praekondition: Ausdr faengt an mit Linksklammer '(', Zahl oder Leerzeichen.
// Z.B. "1" -> 1 und "(1 + 1)" -> 1 + 1
// Im Ausdr koenen verschiedene Ebene von Klammer geben: 0( 1( 2( ... ) ) )
func ParseGrundAusdruck(r *bufio.Reader) (Ausdr, error) {
	symbol, err := leseRune(r)
	if err != nil {
		return nil, err
	}

	if unicode.IsDigit(symbol) { // parse Zahl
		r.UnreadRune() // TODO: better way than unread?
		return ParseZahl(r)

	} else if symbol == '(' { // parse Ausdruck zwischen Klammern
		a, err := ParseTeilAusdruck(r, 0) // fang an mit Rekursionsebene 0
		if err != nil {
			return nil, err
		}

		wantRune := ')' // erwarte schließende Klammer
		gotRune, err := leseRune(r)
		if err != nil || gotRune != wantRune {
			return nil, fmt.Errorf("beim Parsen Aussenausdruk: erwartet ')', gefunden %q", gotRune)
		}

		return a, nil
	} else {
		return nil, fmt.Errorf("unerwartetes Symbol %q", symbol)
	}
}

// ParseZahl liest und verarbeitet eine Zahl aus der Zeichenkette.
func ParseZahl(r *bufio.Reader) (Ausdr, error) {
	var strZahl StrZahl
	for {
		symbol, err := leseRune(r)
		if err == io.EOF {
			break // Dateiende
		}

		if !unicode.IsDigit(symbol) {
			r.UnreadRune() // Keine Ziffer mehr, Symbol zurücklegen
			break
		}
		strZahl.Hinzu(symbol)
	}

	if strZahl.Leer() {
		return nil, fmt.Errorf("Fehler beim Parsen einer Zahl")
	}

	w, err := strZahl.Wert()
	if err != nil {
		return nil, err
	}

	return Ganzezahl{w}, nil
}

// leseRune liest die naechste Rune aus der Zeichenkette und ignoriert die Leerzeichen.
func leseRune(r *bufio.Reader) (rune, error) {
	for {
		symbol, _, err := r.ReadRune()
		if err != nil {
			return 0, err
		}
		if unicode.IsSpace(symbol) {
			continue
		}

		return symbol, nil
	}
}

// OpPrio bestimmt die Prioritaet der Operationen
func OpPrio(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	default:
		return 0
	}
}