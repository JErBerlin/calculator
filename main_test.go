package main

import (
	"strings"
	"testing"
)

func TestNurEineZahl(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Positive: Einfache Zahl", "42", 42, false},
		{"Negative: Keine Zahl", "xyz", 0, true},
	}

	for _, tt := range tests {
		runTestBaueAusdr(t, tt.name, tt.input, tt.want, tt.wantError)
	}
}

func TestZahlInKlammern(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Positive: Zahl in Klammern", "(42)", 42, false},
		{"Negative: Fehlende schließende Klammer", "(42", 0, true},
	}

	for _, tt := range tests {
		runTestBaueAusdr(t, tt.name, tt.input, tt.want, tt.wantError)
	}
}

func TestSummeOderSubtraktion(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Positive: Einfache Summe", "1 + 2", 3, false},
		{"Positive: Einfache Subtraktion", "5 - 3", 2, false},
		{"Negative: Ungültiger Operator", "1 $ 2", 0, true},
	}

	for _, tt := range tests {
		runTestBaueAusdr(t, tt.name, tt.input, tt.want, tt.wantError)
	}
}

func TestMultiplikationOderDivision(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Positive: Einfache Multiplikation", "4 * 2", 8, false},
		{"Positive: Einfache Division", "8 / 2", 4, false},
		{"Negative: Division durch Null", "8 / 0", 0, true},
		{"Negative: Ungültiger Operator", "4 # 2", 0, true},
	}

	for _, tt := range tests {
		runTestBaueAusdr(t, tt.name, tt.input, tt.want, tt.wantError)
	}
}


func TestKomplexeAusdrucke(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Komplexer Ausdruck 1", "3 + 4 * 2", 11, false},
		{"Komplexer Ausdruck 2", "(3 + 4) * 2", 14, false},
		{"Komplexer Ausdruck 3", "3 * (4 + 2) / 2", 9, false},
	}

	for _, tt := range tests {
		runTestBaueAusdr(t, tt.name, tt.input, tt.want, tt.wantError)
	}
}

func TestOperatorVorrang(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Vorrang 1", "1 + 2 * 3", 7, false},
		{"Vorrang 2", "1 * 2 + 3", 5, false},
	}

	for _, tt := range tests {
		runTestBaueAusdr(t, tt.name, tt.input, tt.want, tt.wantError)
	}
}

func TestUngultigeAusdrucke(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Num
		wantError bool
	}{
		{"Leerer Ausdruck", "", 0, true},
		{"Unvollständiger Ausdruck", "42 + ", 0, true},
		{"Nur Operatoren", "+ -", 0, true},
	}

	for _, tt := range tests {
		runTestBaueAusdr(t, tt.name, tt.input, tt.want, tt.wantError)
	}
}

func runTestBaueAusdr(t *testing.T, name, input string, want Num, wantError bool) {
	t.Run(name, func(t *testing.T) {
		rd := strings.NewReader(input)
		a, err := baueAusdr(rd)

		var wert Num
		var errWert error
		if err == nil {
			wert, errWert = a.Wert()
		}
		gotError := (err != nil || errWert != nil)

		if gotError != wantError {
			t.Errorf("Erwarteter Error = %v, erhalten = %v", wantError, gotError)
			return
		}

		if !gotError && wert != want {
			t.Errorf("Erwarteter Wert = %d, erhalten = %d", want, wert)
		}
	})
}
