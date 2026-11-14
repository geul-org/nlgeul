//ff:func feature=token type=helper control=sequence
//ff:what TestLettersOf — the letters of the sentence without its spaces (tabs and wide spaces too), which is what a Span counts
package token

import "testing"

func TestLettersOf(t *testing.T) {
	if got := string(lettersOf(" 아버지 가방에\t들어가신다　")); got != "아버지가방에들어가신다" {
		t.Fatal(got)
	}
	if got := lettersOf(""); len(got) != 0 {
		t.Fatal(got)
	}
}
