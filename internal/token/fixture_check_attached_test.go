//ff:func feature=token type=helper control=iteration dimension=1
//ff:what checkAttached — the candidates of a unit as markers written apart from their word are those given, in order, each a word with no stem that attaches to the unit before, and no higher than a new word
package token

import "testing"

func checkAttached(t *testing.T, in string, wants []string) {
	sc := scanOf(in)
	got := sc.attached()
	if len(got) != len(wants) || sc.attach {
		t.Errorf("%s: %v", in, got)
		return
	}
	for i, p := range got {
		if p.sig != wants[i] || !p.last.word.Attach || p.last.word.Parts[0].Stem || p.verdict > sc.ix.unknown || p.last.end != len([]rune(in)) {
			t.Errorf("%s: %v", in, p.sig)
		}
	}
}
