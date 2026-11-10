//ff:func feature=token type=helper control=sequence
//ff:what TestNativeTokens — a one alone and after 열, the numerals that stand alone, and texts that are none
package token

import "testing"

func TestNativeTokens(t *testing.T) {
	if nativeTokens("한") != 1 || nativeTokens("다섯") != 1 || nativeTokens("열") != 1 || nativeTokens("여러") != 1 || nativeTokens("열두") != 2 {
		t.Fatal("numerals")
	}
	if nativeTokens("스무한") != 0 || nativeTokens("열몇") != 0 || nativeTokens("곰") != 0 || nativeTokens("") != 0 {
		t.Fatal("others")
	}
}
