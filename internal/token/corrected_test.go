//ff:func feature=token type=helper control=sequence
//ff:what TestCorrected — a part written right keeps its readings as they are; a part written wrong gets the sign, 오타정정: and the written form after its readings, and both the sign and the name come from the two constants
package token

import "testing"

func TestCorrected(t *testing.T) {
	if got := corrected("열쇠:동사", ""); got != "열쇠:동사" {
		t.Fatal(got)
	}
	if got := corrected("열쇠:동사", "맞기"); got != "열쇠:동사,오타정정:맞기" {
		t.Fatal(got)
	}
	if got := corrected("열쇠:동사|형용사", "맞기"); got != "열쇠:동사|형용사"+typoSign+typoName+":맞기" {
		t.Fatal(got)
	}
	if typoSign != "," || typoName != "오타정정" {
		t.Fatal(typoSign, typoName)
	}
}
