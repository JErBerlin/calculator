// main_test.go
package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestEinfacheArithmetik(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		wantError bool
	}{
		{"Einfache Summe", "1 + 2", 3, false},
		{"Einfache Subtraktion", "5 - 3", 2, false},
		{"Einfache Multiplikation", "4 * 2", 8, false},
		{"Einfache Division", "8 / 2", 4, false},
		{"Division durch Null", "8 / 0", 0, true},
	}
	runTests(t, tests)
}

func TestEinzelneZahlenUndFehler(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		wantError bool
	}{
		{"Nur eine Zahl", "42", 42, false},
		{"Ungültiger Ausdruck", "abc", 0, true},
	}
	runTests(t, tests)
}

func TestErweiterteArithmetikMitLeerzeichen(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		wantError bool
	}{
		{"Summe mit Leerzeichen", " 1 + 2 ", 3, false},
		{"Subtraktion mit Neuzeile", "\n5\n-\n3\n", 2, false},
		{"Multiplikation mit gemischten Leerzeichen", " 4\t*\t2 ", 8, false},
		{"Division mit Leerzeichen und Neuzeile", " 8 / 2 ", 4, false},
	}
	runTests(t, tests)
}

func TestFehlerhafteEingaben(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		wantError bool
	}{
		{"Verknüpfung mit nur einem Operanden", "42 +", 0, true},
		{"Leerzeichen und Neuzeile gemischt", " \n 42 ", 42, false},
		{"Ungültiger Ausdruck mit Leerzeichen", "a b c", 0, true},
	}
	runTests(t, tests)
}

func TestKomplexeAusdruecke(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		wantError bool
	}{
		// {"Operatoren verschiedener Prioritaet", "3 + 1 * 2 + 2 * 3", 11, false},
		{"Ausdruck mit klammern", "3 + (1 + 2)", 9, false},
		// {"Verschachtelte Klammern", "(1 + (2 * 3)) * 2", 14, false},
	}
	runTests(t, tests)
}

// runTests führt eine Reihe von Tests aus, die in einer Teststruktur definiert sind.
func runTests(t *testing.T, tests []struct {
	name      string
	input     string
	want      int
	wantError bool
}) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rd := strings.NewReader(tt.input)
			reader := bufio.NewReader(rd)
			a, err := ParseAusdruck(reader)

			var wert int
			var errWert error
			if err == nil {
				wert, errWert = a.Wert()
			}
			gotError := (err != nil || errWert != nil)

			if gotError != tt.wantError {
				t.Errorf("Erwarteter Error = %v, erhalten = %v", tt.wantError, gotError)
				return
			}

			if !gotError && wert != tt.want {
				t.Errorf("Erwarteter Wert = %d, erhalten = %d", tt.want, wert)
			}
		})
	}
}
