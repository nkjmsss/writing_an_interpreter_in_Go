package token

type Tokentype string

type Token struct {
	Type    Tokentype
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Tokentype{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) Tokentype {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
