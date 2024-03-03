package main

import (
	"errors"
	"fmt"
	"strings"
)

// ErrUngultigeOperation zeigt an, dass die Operation nicht erkannt wird.
var ErrUngultigeOperation = errors.New("ungültige Operation")

// Verkn definiert ein Interface für arithmetische Verknüpfungen.
type Verkn interface {
	Wert() (Num, error) // Gibt das Ergebnis der Verknüpfung zurück.
	Operation() rune    // Gibt das Symbol der Operation zurück.
	A() Num             // Gibt den Wert von Operand A zurück.
	B() Num             // Gibt den Wert von Operand B zurück.
	SetA(Num)           // Setzt den Wert von Operand A.
	SetB(Num)           // Setzt den Wert von Operand B.
}

// Binop repräsentiert eine binäre Operation (Addition, Subtraktion, Multiplikation, Division).
type Binop struct {
	a, b Num
	op   rune
}

// SetA setzt den Wert von Operand A.
func (p *Binop) SetA(a Num) {
	p.a = a
}

// SetB setzt den Wert von Operand B.
func (p *Binop) SetB(b Num) {
	p.b = b
}

// Wert führt die binäre Operation aus und gibt das Ergebnis zurück.
func (p Binop) Wert() (Num, error) {
	switch p.op {
	case '+':
		return p.a + p.b, nil
	case '-':
		return p.a - p.b, nil
	case '*':
		return p.a * p.b, nil
	case '/':
		if p.b == 0 {
			return 0, fmt.Errorf("Teilen durch Null")
		}
		return p.a / p.b, nil
	default:
		return 0, fmt.Errorf("Ungültige Operation %c", p.op)
	}
}

// Operation gibt das Symbol der Operation zurück.
func (p Binop) Operation() rune {
	return p.op
}

// A gibt den Wert von Operand A zurück.
func (p Binop) A() Num {
	return p.a
}

// B gibt den Wert von Operand B zurück.
func (p Binop) B() Num {
	return p.b
}

// NeueBinop erzeugt eine neue Binop-Instanz mit der spezifizierten Operation.
// Die Operanden werden spaeter mit set-Methoden gesetzt
func NeueBinop(op rune) (*Binop, error) {
	// Unterstuetzte Ops sind +,-,*,/
	if !strings.ContainsRune("+-*/", op) {
		return nil, ErrUngultigeOperation
	}

	return &Binop{
		op: op,
	}, nil
}
