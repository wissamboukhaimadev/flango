package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"
	STRING = "STRING"
	// Operatorspa
	ASSIGN      = "="
	PLUS        = "+"
	DOUBLEPLUS  = "++"
	MINUS       = "-"
	DOUBLEMINUS = "--"
	BANG        = "!"
	ASTERISK    = "*"
	SLASH       = "/"
	LT          = "<"
	GT          = ">"
	INC_STEP    = "+="
	DEC_STEP    = "-="
	MULT_STEP   = "*="
	DIV_STEP    = "/="

	// logical operators
	EQ     = "=="
	NOT_EQ = "!="
	// Delimiters
	COMMA     = ","
	SEMICOLON = "$"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "MAYBE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
	// additional
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"maybe":  LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
