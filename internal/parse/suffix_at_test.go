//ff:func feature=parse type=rule control=sequence
//ff:what TestSuffixAt — the start of a suffix that is there; false for one that is not or is longer than the text
package parse

import "testing"

func TestSuffixAt(t *testing.T) {
	s := []rune("학교에서는")
	if start, ok := suffixAt(s, 4, "에서"); !ok || start != 2 {
		t.Fatal(start, ok)
	}
	if _, ok := suffixAt(s, 5, "에서"); ok {
		t.Fatal("에서 at the end")
	}
	if _, ok := suffixAt(s, 1, "에서"); ok {
		t.Fatal("longer than the text")
	}
	if start, ok := suffixAt(s, 3, ""); !ok || start != 3 {
		t.Fatal("empty form")
	}
}
