package token

import (
	"fmt"
	"testing"
)

func LookupIdent(ident string) TokenType {
	fmt.Println(ident)
	if tok, ok := keywords[ident]; ok {
		fmt.Println(tok, ok)
		return tok
	}
	fmt.Println(IDENT)
	return IDENT
}

func TestLookupIdent(t *testing.T) {
	// lex := LookupIdent("hello")
	mex := LookupIdent("let")
	vex := LookupIdent("fn")
	// aex := LookupIdent("ASSIGN")
	// bex := LookupIdent("=")
	// fmt.Println(lex)
	fmt.Println(mex)
	fmt.Println(vex)
	// fmt.Println(aex)
	// fmt.Println(bex)
}
