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
	LET		= "LET"
	TRUE	= "TRUE"
	FALSE   = "FALSE"
	IF      = "IF"
	ELSE    = "ELSE"
	RETURN  = "RETURN"

	// Operators
	MINUS	= "-"
	BANG	= "!"
	ASTERISK= "*"
	SLASH	= "/"
	EQ		= "=="
	N_EQ	= "!="

	LT = "<"
	GT = ">"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
