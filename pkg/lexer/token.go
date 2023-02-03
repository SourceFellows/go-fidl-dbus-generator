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
	ASTERISK             // *
	COMMA                // ,
	QUOTE                // "
	SINGLE_QUOTE         // '
	CIRCUMFLEX           // ^
	SQUARE_BRACKET_OPEN  // [
	SQUARE_BRACKET_CLOSE // ]
	CURLY_BRACKET_OPEN   // {
	CURLY_BRACKET_CLOSE  // }

	// Keywords
	PACKAGE
	INTERFACE
	IMPORT
	VERSION
	MAJOR
	MINOR
	ATTRIBUTE
	METHOD
	BROADCAST
	STRUCT
	TYPEDEF
	ARRAYDEF
	SELECTIVE
	IN
	OUT
	FIRE_AND_FORGET
)
