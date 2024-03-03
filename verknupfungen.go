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
	A() Ausdr           // Gibt den Operand A zurück.
	B() Ausdr           // Gibt den Operand B zurück.
	SetA(Ausdr)         // Setzt den Wert von Operand A.
	SetB(Ausdr)         // Setzt den Wert von Operand B.
}

// Binop repräsentiert eine binäre Operation (Addition, Subtraktion, Multiplikation, Division).
type Binop struct {
	a, b Ausdr
	op   rune
}

func (p Binop) String() string {
	return fmt.Sprintf("%v %v %v", p.a, p.op, p.b)
}

// SetA setzt den Wert von Operand A.
func (p *Binop) SetA(a Ausdr) {
	p.a = a
}

// SetB setzt den Wert von Operand B.
func (p *Binop) SetB(b Ausdr) {
	p.b = b
}

// Wert führt die binäre Operation aus und gibt das Ergebnis zurück.
func (p Binop) Wert() (Num, error) {
	a, err := p.a.Wert()
	if err != nil {
		return 0, fmt.Errorf("Auswertung von A in Verknupfung %s fehlgeschlagen: %w", p, err)
	}
	b, err := p.b.Wert()
	if err != nil {
		return 0, fmt.Errorf("Auswertung von B in Verknupfung %s fehlgeschlagen: %w", p, err)
	}

	switch p.op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, fmt.Errorf("Teilen durch Null")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Ungültige Operation %c", p.op)
	}
}

// Operation gibt das Symbol der Operation zurück.
func (p Binop) Operation() rune {
	return p.op
}

// A gibt den Wert von Operand A zurück.
func (p Binop) A() Ausdr {
	return p.a
}

// B gibt den Wert von Operand B zurück.
func (p Binop) B() Ausdr {
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
