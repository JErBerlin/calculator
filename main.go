package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	strAusdr := "abc"
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

// Ausdr repräsentiert einen Ausdruck, der ausgewertet werden kann.
type Ausdr interface {
	Wert() (int, error)
}

// Ganzezahl ist der (numerischer) Ausdruck einer Ganzezahl
type Ganzezahl struct {
	n int
}

// Wert gibt den numersichen Wert der Ganzezahl zurück
func (z Ganzezahl) Wert() (int, error) {
	return z.n, nil
}

// Verkn definiert ein Interface für arithmetische Verknüpfungen.
type Verkn interface {
	Wert() (int, error) // Gibt das Ergebnis der Verknüpfung zurück.
	Operation() rune    // Gibt das Symbol der Operation zurück.
	A() int             // Gibt den Wert von Operand A zurück.
	B() int             // Gibt den Wert von Operand B zurück.
	SetA(int)           // Setzt den Wert von Operand A.
	SetB(int)           // Setzt den Wert von Operand B.
}

// Summe repräsentiert eine Addition.
type Summe struct {
	a, b int
}

func (s *Summe) SetA(a int) {
	s.a = a
}

func (s *Summe) SetB(b int) {
	s.b = b
}

func (s Summe) Wert() (int, error) {
	return s.a + s.b, nil
}

func (s Summe) Operation() rune {
	return '+'
}

func (s Summe) A() int {
	return s.a
}

func (s Summe) B() int {
	return s.b
}

// Reste repräsentiert eine Subtraktion.
type Reste struct {
	a, b int
}

func (r *Reste) SetA(a int) {
	r.a = a
}

func (r *Reste) SetB(b int) {
	r.b = b
}

func (r Reste) Wert() (int, error) {
	return r.a - r.b, nil
}

func (r Reste) Operation() rune {
	return '-'
}

func (r Reste) A() int {
	return r.a
}

func (r Reste) B() int {
	return r.b
}

// Mult repräsentiert eine Multiplikation.
type Mult struct {
	a, b int
}

func (m *Mult) SetA(a int) {
	m.a = a
}

func (m *Mult) SetB(b int) {
	m.b = b
}

func (m Mult) Wert() (int, error) {
	return m.a * m.b, nil
}

func (m Mult) Operation() rune {
	return '*'
}

func (m Mult) A() int {
	return m.a
}

func (m Mult) B() int {
	return m.b
}

// Div repräsentiert eine Division.
type Div struct {
	a, b int
}

func (d *Div) SetA(a int) {
	d.a = a
}

func (d *Div) SetB(b int) {
	d.b = b
}

func (d Div) Wert() (int, error) {
	if d.b == 0 {
		return 0, fmt.Errorf("Teilen durch Null") // TODO: return a more meaningful value like Inf or an error
	}
	if d.a == 0 {
		return 0, nil
	}
	return d.a / d.b, nil
}

func (d Div) Operation() rune {
	return '/'
}

func (d Div) A() int {
	return d.a
}

func (d Div) B() int {
	return d.b
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

// baueAusdr baut einen Ausdruck aus einer Rune-Zeichenkette im Reader
func baueAusdr(r io.Reader) (Ausdr, error) {
	reader := bufio.NewReader(r)

	var exp Ausdr
	var sz StrZahl // String um die gelesenen Ziffer zu speichern und dann zum Nummerischen Wert zu konvertieren

	for weiterLesen := true; weiterLesen; {
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

		// lese eine Zahl und ggf. eine Verknuepfung
		switch {
		// lese eine Ziffer von Operanden a oder Operanden b
		case unicode.IsDigit(symbol):
			sz.Hinzu(symbol) // fuegt gelesene Ziffer hinzu

		// lese Summe-Verknuepfung
		case symbol == '+':
			s := new(Summe)
			w, err := sz.Wert()
			if err != nil {
				return exp, fmt.Errorf("baueAusdr: vor der Verknüpfung erwarteter Operanden A nicht vorhanden")
			}
			s.SetA(w)  // speichere erstem Operanden A im Ausdruck
			sz.Reset() // reset Zahlenleser, um Operanden b zu lesen
			exp = s    // speichere die Summe als Ausdrueck exp

		// lese Reste-Verknuepfung
		case symbol == '-':
			r := new(Reste)
			w, err := sz.Wert()
			if err != nil {
				return exp, fmt.Errorf("baueAusdr: vor der Verknüpfung erwarteter Operanden A nicht vorhanden")
			}
			r.SetA(w)
			sz.Reset()
			exp = r

		// lese Mult-Verknuepfung
		case symbol == '*':
			m := new(Mult)
			w, err := sz.Wert()
			if err != nil {
				return exp, fmt.Errorf("baueAusdr: vor der Verknüpfung erwarteter Operanden A nicht vorhanden")
			}
			m.SetA(w)
			sz.Reset()
			exp = m

		// lese Div-Verknuepfung
		case symbol == '/':
			d := new(Div)
			w, err := sz.Wert()
			if err != nil {
				return exp, fmt.Errorf("baueAusdr: vor der Verknüpfung erwarteter Operanden A nicht vorhanden")
			}
			d.SetA(w)
			sz.Reset()
			exp = d

		default:
			weiterLesen = false
		}
	}

	w, err := sz.Wert() // Die gelesene Zeichenkette ist eine Zahl alleine oder der zweite Operand b
	if err != nil {     // etwas ist schiefgelaufen: wir brauchen wenigstens eine Zahl
		return exp, fmt.Errorf("Zahlen konnten nicht vollstaendig gelesen werden")
	}

	// wenn der Ausdrueck eine Verknuepfung ist, fehlt noch den zweiten Operanden zu speichern
	if v, ok := exp.(Verkn); ok {
		v.SetB(w)
	} else { // sonst speichern wir eine Zahl alleine
		exp = Ganzezahl{w}
	}

	return exp, nil
}
