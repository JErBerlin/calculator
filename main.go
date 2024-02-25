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
	strAusdr := "42"
	rd := strings.NewReader(strAusdr)

	a, err := baueAusdr(rd)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	if v, ok := a.(Verkn); ok {
		fmt.Printf("Ausdruck ist die Verknuepfung: %d %s %d\n", v.A(), string(v.Operation()), v.B())
		fmt.Printf("Wert(%s) = %d\n", strAusdr, a.Wert())
	} else if _, ok := a.(Ganzezahl); ok {
		fmt.Println("Ausdruck ist eine Ganzezahl")
		fmt.Printf("Wert(%s) = %d\n", strAusdr, a.Wert())
	} else {
		fmt.Println("Ausdruck war unlesbar. Unfug.")
	}
}

type Ausdr interface {
	Wert() int
}

type Ganzezahl struct {
	n int
}

func (z Ganzezahl) Wert() int {
	return z.n
}

// Verkn definiert ein Interface für arithmetische Verknüpfungen.
type Verkn interface {
	Wert() int       // Gibt das Ergebnis der Verknüpfung zurück.
	Operation() rune // Gibt das Symbol der Operation zurück.
	A() int          // Gibt den Wert von Operand A zurück.
	B() int          // Gibt den Wert von Operand B zurück.
	SetA(int)        // Setzt den Wert von Operand A.
	SetB(int)        // Setzt den Wert von Operand B.
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

func (s Summe) Wert() int {
	return s.a + s.b
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

func (r Reste) Wert() int {
	return r.a - r.b
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

func (m Mult) Wert() int {
	return m.a * m.b
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

func (d Div) Wert() int {
	if d.b == 0 {
		return 0 // TODO: return a more meaningful value like Inf or an error
	}
	return d.a / d.b
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
// Beim leeren hinzufuegen bekommt man den numerischen Wert zurueck.
// Beim hinzufuegen von 'R' wird die Zeichenkette geleert.
func (z *StrZahl) Hinzu(d rune) {
	z.zahlAlsString += string(d)
}

// Wert konvertiert die Zeichenkettendarstellung der Zahl in eine Ganzzahl.
func (z *StrZahl) Wert() int {
	i, _ := strconv.Atoi(z.zahlAlsString) // Fehler ignorieren, da immer eine Zahl erwartet wird
	return i
}

func (z *StrZahl) Reset() {
	z.zahlAlsString = ""
}

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
			s.SetA(sz.Wert()) // speichere den bereits gelesenen Operanden a
			sz.Reset()        // reset Zahlenleser, um Operanden b zu lesen
			exp = s           // speichere die Summe als Ausdrueck exp

		// lese Reste-Verknuepfung
		case symbol == '-':
			r := new(Reste)
			r.SetA(sz.Wert()) // speichere den bereits gelesenen Operanden a
			sz.Reset()        // reset Zahlenleser, um Operanden b zu lesen
			exp = r           // speichere die Reste als Ausdrueck exp

		// lese Mult-Verknuepfung
		case symbol == '*':
			m := new(Mult)
			m.SetA(sz.Wert()) // speichere den bereits gelesenen Operanden a
			sz.Reset()        // reset Zahlenleser, um Operanden b zu lesen
			exp = m           // speichere die Mult als Ausdrueck exp

		// lese Div-Verknuepfung
		case symbol == '/':
			d := new(Div)
			d.SetA(sz.Wert()) // speichere den bereits gelesenen Operanden a
			sz.Reset()        // reset Zahlenleser, um Operanden b zu lesen
			exp = d           // speichere die Div als Ausdrueck exp

		default:
			weiterLesen = false
		}
	}

	if sz.zahlAlsString == "" { // etwas ist schiefgelaufen: wir brauchen wenigstens eine Zahl
		return exp, fmt.Errorf("Zahlen konnten nicht vollstaendig gelesen werden")
	}

	// wenn der Ausdrueck eine Verknuepfung ist, fehlt noch den zweiten Operanden zu speichern
	if v, ok := exp.(Verkn); ok {
		v.SetB(sz.Wert())
	} else {
		exp = Ganzezahl{sz.Wert()}
	}

	return exp, nil
}
