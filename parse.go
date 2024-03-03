package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// ParseAusdruck parses and evaluates expressions with addition and subtraction.
func ParseAusdruck(scanner *SymbolScanner) (Ausdr, error) {
	// Parse the first term
	lhs, err := ParseTerm(scanner)
	if err != nil {
		return nil, err
	}

	for {
		// Look at the next token without consuming it
		op := scanner.Scanner.Peek()
		if op != '+' && op != '-' {
			break
		}
		scanner.Scanner.Scan() // Consume the operator

		// Parse the next term
		rhs, err := ParseTerm(scanner)
		if err != nil {
			return nil, err
		}

		// Create a Binop with the lhs and rhs
		binOp, err := NeueBinop(scanner.LastToken)
		if err != nil {
			return nil, err
		}
		binOp.SetA(lhs)
		binOp.SetB(rhs)

		// The current binOp becomes the new lhs for any further additions/subtractions
		lhs = binOp
	}

	return lhs, nil
}

// ParseTerm parses terms that involve multiplication and division.
func ParseTerm(scanner *SymbolScanner) (Ausdr, error) {
	lhs, err := ParseFaktor(scanner)
	if err != nil {
		return nil, err
	}

	for {
		op := scanner.Scanner.Peek()
		if op != '*' && op != '/' {
			break
		}
		scanner.Scanner.Scan() // Consume the operator

		rhs, err := ParseFaktor(scanner)
		if err != nil {
			return nil, err
		}

		// Create a Binop with the lhs and rhs
		binOp, err := NeueBinop(scanner.LastToken)
		if err != nil {
			return nil, err
		}
		binOp.SetA(lhs)
		binOp.SetB(rhs)

		// The current binOp becomes the new lhs for any further multiplications/divisions
		lhs = binOp
	}

	return lhs, nil
}

// ParseFaktor parses a factor, which can be a number or a parenthesized expression.
func ParseFaktor(scanner *SymbolScanner) (Ausdr, error) {
	tok := scanner.Scanner.TokenText()
	// Define a regular expression to match integer numbers
	intRegex := regexp.MustCompile(`^-?\d+$`)

	if intRegex.MatchString(tok) {
		// If the token matches the regex, it's an integer
		zahl, err := strconv.Atoi(tok)
		if err != nil {
			return nil, fmt.Errorf("Fehler beim Parsen der Zahl: %w", err)
		}
		return Zahl{Num(zahl)}, nil
	} else if tok == "(" {
		// Handle parentheses by recursively parsing the expression inside
		ausdruck, err := ParseAusdruck(scanner)
		if err != nil {
			return nil, err
		}
		if scanner.Scanner.Scan(); scanner.Scanner.TokenText() != ")" {
			return nil, fmt.Errorf("Erwartete schließende Klammer")
		}
		return ausdruck, nil
	} else {
		// If the token is not an integer or an opening parenthesis, it's unexpected
		return nil, fmt.Errorf("Unerwartetes Token: %s", tok)
	}
}
