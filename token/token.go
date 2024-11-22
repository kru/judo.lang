package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + literal
	IDENT = "IDENT"

	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"

	// Delimiters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
