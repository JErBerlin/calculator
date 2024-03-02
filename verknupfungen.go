package main

import (
	"errors"
	"fmt"
)

// ErrInvalidOperation zeigt an, dass die Operation nicht erkannt wird.
var ErrInvalidOperation = errors.New("ungültige Operation")

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

func leseVerknupfung(symbol rune) (Verkn, error) {
	switch symbol {
	case '+':
		return new(Summe), nil

	case '-':
		return new(Reste), nil

	case '*':
		return new(Mult), nil

	case '/':
		return new(Div), nil

	default:
	}
	return nil, ErrInvalidOperation
}
