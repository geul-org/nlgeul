//ff:func feature=parse type=reader control=sequence
//ff:what TestStemTokens — a path read with the Hangul after its last slash as its name is one token, as is a path with no head and one slash; the same letters as a new word, and any other stem, are as many tokens as tokens counts
package parse

import "testing"

func TestStemTokens(t *testing.T) {
	sc := scanOf("~/a/b/회의록")
	sc.search()
	path := stemCandidate{kindCandidate{0, pathName, 1}, "~/a/b/회의록", ""}
	if got := sc.stemTokens(0, 9, path); got != 1 {
		t.Fatal(got)
	}
	if got := sc.stemTokens(0, 9, stemCandidate{kindCandidate{}, "~/a/b/회의록", ""}); got != sc.tokens(0, 9) || got < 2 {
		t.Fatal(got)
	}
	if got := sc.stemTokens(0, 6, path); got != 1 || got != sc.tokens(0, 6) {
		t.Fatal("the path up to its slash is an atom", got)
	}
	if got := sc.stemTokens(6, 9, stemCandidate{kindCandidate{0, "열쇠:체언", 1}, "회의록", ""}); got != 3 {
		t.Fatal(got)
	}
}
