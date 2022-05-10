package lexer

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	// Literals
	IDENT
	DESCRIPTION

	// Misc characters
	ASTERISK     // *
	COMMA        // ,
	QUOTE        // "
	SINGLE_QUOTE // '

	// Keywords
	PACKAGE
	INTERFACE
	VERSION
	ATTRIBUTE
	METHOD
	BROADCAST
	STRUCT
	IN
	OUT
)
