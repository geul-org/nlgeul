//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestHandNamed — a sign a hand-written table reads is named by hand: a bracket (the table of pairs), a comma (the sign table), % (the unit table); a sign no such table reads (○, ⓐ) is not
package parse

import "testing"

func TestHandNamed(t *testing.T) {
	ix := &index{}
	ix.addDots()
	ix.addSigns()
	ix.formEntries = map[string]*entry{}
	for r, want := range map[rune]bool{'(': true, ',': true, '%': true, '○': false, 'ⓐ': false} {
		if got := ix.handNamed(r); got != want {
			t.Errorf("handNamed(%c) = %v", r, got)
		}
	}
}
