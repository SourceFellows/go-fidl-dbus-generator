package lexer

import (
	"fmt"
	"io"
)

// Fidl represents a FIDL file.
type (
	Fidl struct {
		PackageInfo   *PackageInfo
		InterfaceInfo *InterfaceInfo
		Attributes    []Attribute
		Methods       []Method
		Broadcasts    []Broadcast
		Structs       []Struct
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
		IsArray     bool
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
		PackageInfo:   nil,
		InterfaceInfo: nil,
		Attributes:    nil,
		Methods:       nil,
		Broadcasts:    nil,
		Structs:       nil,
	}

	packageInfo, err := p.scanPackageInfo()
	if err != nil {
		return nil, err
	}

	fidl.PackageInfo = packageInfo

	interfaceInfo, err := p.scanInterfaceInfo()
	if err != nil {
		return nil, err
	}

	fidl.InterfaceInfo = interfaceInfo

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

			attr := p.scanAttribute()
			attr.Description = description

			fidl.Attributes = append(fidl.Attributes, attr)
		case METHOD:
			if fidl.Methods == nil {
				fidl.Methods = []Method{}
			}

			meth := p.scanMethod()
			meth.Description = description

			fidl.Methods = append(fidl.Methods, meth)
		case BROADCAST:
			if fidl.Broadcasts == nil {
				fidl.Broadcasts = []Broadcast{}
			}

			bc := p.scanBroadcast()
			bc.Description = description

			fidl.Broadcasts = append(fidl.Broadcasts, bc)
		case STRUCT:
			if fidl.Structs == nil {
				fidl.Structs = []Struct{}
			}

			str := p.scanStruct()
			str.Description = description

			fidl.Structs = append(fidl.Structs, str)
		default:
			// ignore unknown input for now
			continue

		}
	}

	// Return the successfully parsed FIDL.
	return fidl, nil
}

func (p *Parser) scanPackageInfo() (*PackageInfo, error) {
	packageInfo := &PackageInfo{}

	// First token should be a "package" keyword.
	tok, lit := p.scanIgnoreWhitespace()
	if tok != PACKAGE {
		return nil, fmt.Errorf("found %q, expected package", lit)
	}

	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("expected IDENT, got %v", tok)
	}

	packageInfo.Name = lit

	// TODO: scan optional imports

	return packageInfo, nil
}

func (p *Parser) scanInterfaceInfo() (*InterfaceInfo, error) {
	interfaceInfo := &InterfaceInfo{}

	tok, lit := p.scanIgnoreWhitespace()
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

	interfaceInfo.Description = desc
	interfaceInfo.Name = lit

	// TODO: scan interface version

	return interfaceInfo, nil
}

func (p *Parser) scanAttribute() Attribute {
	attr := Attribute{}
	_, lit := p.scanIgnoreWhitespace()
	attr.Type = lit
	tok, lit := p.scanIgnoreWhitespace()
	if tok == SQUARE_BRACKET_OPEN {
		attr.IsArray = true
		// ignore SQUARE_BRACKET_CLOSE
		p.scanIgnoreWhitespace()

		// scan param name
		_, lit = p.scanIgnoreWhitespace()
		attr.Name = lit
	} else {
		attr.Name = lit
	}

	return attr
}

func (p *Parser) scanMethod() Method {
	meth := Method{}
	_, lit := p.scanIgnoreWhitespace()
	meth.Name = lit
	meth.In, meth.Out = p.scanParams()

	return meth
}

func (p *Parser) scanBroadcast() Broadcast {
	bc := Broadcast{}
	_, lit := p.scanIgnoreWhitespace()
	bc.Name = lit
	_, bc.Out = p.scanParams()

	return bc
}

func (p *Parser) scanStruct() Struct {
	str := Struct{}
	_, lit := p.scanIgnoreWhitespace()
	str.Name = lit

	str.Fields = p.scanStructParams()

	return str
}

func (p *Parser) scanParam() Param {
	param := Param{}

	tok, lit := p.scanIgnoreWhitespace()
	if tok == DESCRIPTION {
		param.Description = lit
		// scan param type
		_, lit = p.scanIgnoreWhitespace()
	}

	param.Type = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok == SQUARE_BRACKET_OPEN {
		param.IsArray = true
		// ignore SQUARE_BRACKET_CLOSE
		p.scanIgnoreWhitespace()

		// scan param name
		tok, lit = p.scanIgnoreWhitespace()
		if tok == CIRCUMFLEX {
			tok, lit = p.scanIgnoreWhitespace()
			param.Name = fmt.Sprintf("^%s", lit)
		} else {
			param.Name = lit
		}

	} else {
		if tok == CIRCUMFLEX {
			tok, lit = p.scanIgnoreWhitespace()
			param.Name = fmt.Sprintf("^%s", lit)
		} else {
			param.Name = lit
		}
	}

	return param
}

func (p *Parser) scanParams() ([]Param, []Param) {
	var inParams []Param
	var outParams []Param
	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == CURLY_BRACKET_OPEN {
			continue
		}

		if tok == IN {
			for {
				tok, _ = p.scanIgnoreWhitespace()
				if tok == CURLY_BRACKET_OPEN {
					continue
				}

				if tok == CURLY_BRACKET_CLOSE {
					break
				}

				p.unscan()
				param := p.scanParam()
				inParams = append(inParams, param)
			}

			continue
		}

		if tok == OUT {
			for {
				tok, _ = p.scanIgnoreWhitespace()
				if tok == CURLY_BRACKET_OPEN {
					continue
				}

				if tok == CURLY_BRACKET_CLOSE {
					break
				}

				p.unscan()
				param := p.scanParam()
				outParams = append(outParams, param)
			}

			continue
		}

		break
	}

	return inParams, outParams
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

func (p *Parser) scanStructParams() []Param {
	var params []Param

	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == CURLY_BRACKET_OPEN {
			continue
		}

		if tok == CURLY_BRACKET_CLOSE {
			break
		}

		p.unscan()
		param := p.scanParam()
		params = append(params, param)
	}

	return params
}
