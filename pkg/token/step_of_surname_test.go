//ff:func feature=token type=reader control=sequence
//ff:what TestStepOfSurname — a family name read before a title is valued by its syllables (one token: log10 of 1/11172, less the cost of where it stands), not by the word table's count of its letters, which a known word of the same letters is valued by; its verdict is 1 and its part is named 성:김
package token

import (
	"math"
	"testing"
)

func TestStepOfSurname(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-surname.tsv", Spacing: []string{"48"}})
	if err != nil {
		t.Fatal(err)
	}
	sc := &scan{ix: ix, s: []rune("김군은")}
	sc.prepare()
	at := tail{start: 1}
	sur := sc.stepOf(0, 1, at, stemCandidate{kindCandidate{kindNominal, "성:김", 1}, "김", ""})
	want := math.Log10(1.0/syllables) - ix.structure(0, nil)
	if math.Abs(sur.pref-want) > 1e-9 || sur.verdict != 1 || sur.word.Parts[0].Name != (Name{"성", "김"}) || sur.kind != kindNominal {
		t.Fatal(sur.pref, want, sur.verdict, sur.word.Parts[0].Name)
	}
	known := sc.stepOf(0, 1, at, stemCandidate{kindCandidate{kindNominal, "열쇠:체언", 1}, "김", ""})
	counted := math.Log10(float64(ix.taken("김", tailClass(nil), kindNominal))/ix.lexTotal) - ix.structure(0, nil)
	if ix.taken("김", tailClass(nil), kindNominal) == 0 || math.Abs(known.pref-counted) > 1e-9 || math.Abs(known.pref-sur.pref) < 1e-6 {
		t.Fatal("a known word of the same letters is valued by its count", known.pref, counted, sur.pref)
	}
}
