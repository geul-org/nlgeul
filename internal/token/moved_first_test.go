//ff:func feature=token type=helper control=sequence
//ff:what TestMovedFirst — the sign taken off the first word leaves a word with no stem (Attach), the same verdict, the preference less the sign's own and the end counted from after the sign; a word with only the sign gives false
package token

import "testing"

func TestMovedFirst(t *testing.T) {
	w := Eojeol{Parts: []Part{{Text: ")", Stem: true}, {Text: "를", Names: []Name{{Kind: "목적격", Form: "를"}}}}, End: 2}
	e := step{end: 2, word: w, verdict: 1, pref: -6, sig: "x"}
	got, ok := movedFirst(e, 1, -5)
	if !ok || got.end != 1 || got.word.End != 1 || !got.word.Attach || len(got.word.Parts) != 1 || got.verdict != 1 || got.pref != -1 || got.sig != "+를[목적격:를]" {
		t.Fatal(got, ok)
	}
	if len(w.Parts) != 2 || w.Attach {
		t.Fatal("changed", w)
	}
	if _, ok = movedFirst(step{word: Eojeol{Parts: []Part{{Text: ")"}}}}, 1, 0); ok {
		t.Fatal("only the sign")
	}
}
