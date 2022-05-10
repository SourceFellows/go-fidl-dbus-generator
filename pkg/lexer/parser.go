package lexer

import (
	"fmt"
	"io"
)

// Fidl represents a FIDL file.
type (
	Fidl struct {
		Package    PackageInfo
		Interface  InterfaceInfo
		Attributes []Attribute
		Methods    []Method
		Broadcasts []Broadcast
		Structs    []Struct
	}

	PackageInfo struct {
		Name    string
		Imports []string
	}

	InterfaceInfo struct {
		Name         string
		Description  string
		MajorVersion int
		MinorVersion int
	}

	Attribute struct {
		Description string
		Type        string
		Name        string
	}

	Method struct {
		Description string
		Name        string
		In          []Param
		Out         []Param
	}

	Broadcast struct {
		Description string
		Name        string
		Out         []Param
	}

	Param struct {
		Description string
		Type        string
		Name        string
		IsArray     bool
	}

	Struct struct {
		Description string
		Name        string
		Fields      []Param
	}
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parses a FIDL file.
func (p *Parser) Parse() (*Fidl, error) {
	fidl := &Fidl{
		Package:    PackageInfo{},
		Interface:  InterfaceInfo{},
		Attributes: nil,
		Methods:    nil,
		Broadcasts: nil,
		Structs:    nil,
	}

	// First token should be a "SELECT" keyword.
	tok, lit := p.scanIgnoreWhitespace()
	if tok != PACKAGE {
		return nil, fmt.Errorf("found %q, expected package", lit)
	}

	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("expected IDENT, got %v", tok)
	}

	fidl.Package.Name = lit

	tok, lit = p.scanIgnoreWhitespace()
	var desc string
	if tok == DESCRIPTION {
		desc = lit
		// ignore keyword "interface" after description
		p.scanIgnoreWhitespace()
	}

	if tok != DESCRIPTION && tok != INTERFACE {
		return nil, fmt.Errorf("expected interface definition but got %v: %v", tok, lit)
	}

	tok, lit = p.scanIgnoreWhitespace()

	fidl.Interface.Description = desc
	fidl.Interface.Name = lit

	// TODO: scan interface version

	for {
		innerTok, innerLit := p.scanIgnoreWhitespace()
		if innerTok == EOF {
			break
		}

		var description string
		if innerTok == DESCRIPTION {
			description = innerLit
			innerTok, innerLit = p.scanIgnoreWhitespace()
		}

		switch innerTok {
		case ATTRIBUTE:
			if fidl.Attributes == nil {
				fidl.Attributes = []Attribute{}
			}

			attr := Attribute{
				Description: description,
			}
			_, innerLit = p.scanIgnoreWhitespace()
			attr.Type = innerLit
			_, innerLit = p.scanIgnoreWhitespace()
			attr.Name = innerLit

			fidl.Attributes = append(fidl.Attributes, attr)
		case METHOD:
			if fidl.Methods == nil {
				fidl.Methods = []Method{}
			}

			meth := Method{
				Description: description,
			}
			_, innerLit = p.scanIgnoreWhitespace()
			meth.Name = innerLit

			// TODO parse PARAMS

			fidl.Methods = append(fidl.Methods, meth)
		case BROADCAST:
			if fidl.Broadcasts == nil {
				fidl.Broadcasts = []Broadcast{}
			}

			bc := Broadcast{
				Description: description,
			}
			_, innerLit = p.scanIgnoreWhitespace()
			bc.Name = innerLit

			// TODO parse PARAMS

			fidl.Broadcasts = append(fidl.Broadcasts, bc)
		case STRUCT:
			if fidl.Structs == nil {
				fidl.Structs = []Struct{}
			}

			str := Struct{
				Description: description,
			}
			_, innerLit = p.scanIgnoreWhitespace()
			str.Name = innerLit

			// TODO parse PARAMS

			fidl.Structs = append(fidl.Structs, str)
		default:
			// ignore non known input for now
			continue

		}
	}

	// Return the successfully parsed FIDL.
	return fidl, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WHITESPACE {
		tok, lit = p.scan()
	}

	return
}
