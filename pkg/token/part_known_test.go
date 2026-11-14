//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestPartKnown — a stem no table knows is not known; a stem a table names, a stem written wrong, a marker, a sign with no name and a sign of a pair are known
package token

import "testing"

func TestPartKnown(t *testing.T) {
	if (Part{Text: "쪼간", Stem: true}).Known() {
		t.Fatal("a new word")
	}
	for _, p := range []Part{
		{Text: "물", Stem: true, Name: Name{"열쇠", "체언"}},
		{Text: "맡기", Written: "맞기", Stem: true, Name: Name{"열쇠", "동사"}},
		{Text: "을", Name: Name{"목적격", "을"}},
		{Text: "->", Stem: true},
		{Text: "(", Stem: true, Name: Name{"열기", "1"}, Depth: 1},
	} {
		if !p.Known() {
			t.Fatal(p)
		}
	}
}
