package main

import (
	"strings"
	"testing"
)

func TestBaueAusdr(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Einfache Summe", "1 + 2", 3, false},
		{"Einfache Subtraktion", "5 - 3", 2, false},
		{"Einfache Multiplikation", "4 * 2", 8, false},
		{"Einfache Division", "8 / 2", 4, false},
		{"Division durch Null", "8 / 0", 0, true},
		{"Nur eine Zahl", "42", 42, false},
		{"Ungültiger Ausdruck", "abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rd := strings.NewReader(tt.input)
			a, err := baueAusdr(rd)

			var wert Num
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

func TestBaueAusdrMitErweitertenFaellen(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Summe mit Leerzeichen", " 1 + 2 ", 3, false},
		{"Subtraktion mit Neuzeile", "\n5\n-\n3\n", 2, false},
		{"Multiplikation mit gemischten Leerzeichen", " 4\t*\t2 ", 8, false},
		{"Division mit Leerzeichen und Neuzeile", " 8 / 2 ", 4, false},
		{"Verknüpfung mit nur einem Operanden", "42 +", 0, true},
		{"Leerzeichen und Neuzeile gemischt", " \n 42 ", 42, false},
		{"Ungültiger Ausdruck mit Leerzeichen", "a b c", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rd := strings.NewReader(tt.input)
			a, err := baueAusdr(rd)

			var wert Num
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
