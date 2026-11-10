//ff:func feature=token type=helper control=sequence
//ff:what TestSinoTokens — digits and greater numbers in a run, 수 before a greater number; two digits side by side, 수 alone or inside, and other letters are no numeral
package token

import "testing"

func TestSinoTokens(t *testing.T) {
	if sinoTokens("삼") != 1 || sinoTokens("삼십") != 2 || sinoTokens("이백오십") != 4 || sinoTokens("만") != 1 || sinoTokens("수십") != 2 {
		t.Fatal("numerals")
	}
	if sinoTokens("이삼") != 0 || sinoTokens("수") != 0 || sinoTokens("수삼") != 0 || sinoTokens("십수") != 0 || sinoTokens("삼곰") != 0 || sinoTokens("") != 0 {
		t.Fatal("others")
	}
}
