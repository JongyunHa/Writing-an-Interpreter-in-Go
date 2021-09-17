package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT 식별자 + 리터럴
	IDENT = "IDENT" // add, foobar, x, y ...
	INT   = "INT"   //. 123434566

	// ASSIGN 연산자
	ASSIGN = "="
	PLUS   = "+"

	// COMMA 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
