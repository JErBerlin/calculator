package main

import (
	"io"
	"text/scanner"
)

// SymbolScanner is a struct that encapsulates a scanner.Scanner
// and has the functionality to read specified symbols from an input source.
type SymbolScanner struct {
	Scanner   scanner.Scanner // Embeds the scanner
	LastToken rune            // Holds the last read token
}

// NewSymbolScanner initializes the SymbolScanner with an io.Reader.
// It configures the scanner to recognize integers, arithmetic symbols, and parentheses.
func NewSymbolScanner(r io.Reader) *SymbolScanner {
	s := &SymbolScanner{}
	s.Scanner.Init(r)
	// Configure the scanner to recognize integers and individual characters
	s.Scanner.Mode = scanner.ScanInts | scanner.ScanChars
	return s
}

// ReadToken reads the next token from the input and updates LastToken.
// It returns the token and its text, or scanner.EOF if there are no more tokens.
func (s *SymbolScanner) ReadToken() (rune, string) {
	tok := s.Scanner.Scan()
	s.LastToken = tok
	return tok, s.Scanner.TokenText()
}
