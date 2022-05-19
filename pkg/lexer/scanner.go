package lexer

import (
	"bufio"
	"bytes"
	"io"
)

var eof = rune(0)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if ch == '<' {
		ch2 := s.read()
		if ch2 == '*' {
			// ignore '*' start of description
			s.read()
			return s.scanDescription()
		}
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case '"':
		return QUOTE, string(ch)
	case '\'':
		return SINGLE_QUOTE, string(ch)
	case '^':
		return CIRCUMFLEX, string(ch)
	case '[':
		return SQUARE_BRACKET_OPEN, string(ch)
	case ']':
		return SQUARE_BRACKET_CLOSE, string(ch)
	case '{':
		return CURLY_BRACKET_OPEN, string(ch)
	case '}':
		return CURLY_BRACKET_CLOSE, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		ch := s.read()
		if ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WHITESPACE, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		ch := s.read()
		if ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '.' && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch buf.String() {
	case "package":
		return PACKAGE, buf.String()
	case "import":
		return IMPORT, buf.String()
	case "interface":
		return INTERFACE, buf.String()
	case "version":
		return VERSION, buf.String()
	case "major":
		return MAJOR, buf.String()
	case "minor":
		return MINOR, buf.String()
	case "attribute":
		return ATTRIBUTE, buf.String()
	case "method":
		return METHOD, buf.String()
	case "in":
		return IN, buf.String()
	case "out":
		return OUT, buf.String()
	case "broadcast":
		return BROADCAST, buf.String()
	case "struct":
		return STRUCT, buf.String()
	case "typedef":
		return TYPEDEF, buf.String()
	case "array":
		return ARRAYDEF, buf.String()
	case "selective":
		return SELECTIVE, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

func (s *Scanner) scanDescription() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	for {
		ch := s.read()
		// TODO: use better approach to determine end of description
		if ch != '*' {
			buf.WriteRune(ch)
		} else {
			ch2 := s.read()
			if ch2 == '*' {
				// ignore end of description '>'
				s.read()
				break
			} else {
				buf.WriteRune(ch)
				buf.WriteRune(ch2)
			}
		}
	}

	// return only content of description <** CONTENT **>
	return DESCRIPTION, buf.String()
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
